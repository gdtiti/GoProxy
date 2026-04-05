# webui 模块

`webui` 是 GoProxy 的单体管理入口：它在同一进程里提供登录页、仪表盘 HTML、访客只读 API、管理员写 API，以及订阅添加/贡献接口；配置保存后会立刻通知主流程调整池子参数。

## 模块作用

- 这个模块主要负责：页面渲染、登录态校验、只读/管理员 API 分层、配置保存、订阅入口、日志与代理状态查询。
- 当前边界：
  - `server.go`：HTTP 路由、登录态、REST API
  - `html.go`：登录页嵌入式模板
  - `dashboard.go`：仪表盘嵌入式模板
  - `theme.go`：共享主题 token、主题选择器与前端持久化脚本

## 与项目主页面的关系

- 上游页面：[架构总览](../架构总览.md)、[快速开始](../快速开始.md)
- 下游页面：[custom 模块](./custom.md)、[开放问题](../开放问题.md)
- 是否属于核心路径：是，既是管理入口，也是订阅和配置变更的触发源。

## 入口与装配

- 启动 / 注册入口：
  - `main.go:99-103` 调用 `webui.New(...).Start()`
  - `Start()` 中注册首页、登录、只读 API、管理员 API 和订阅管理 API（证据等级：A）
- 配置入口：
  - 构造函数接收 `cfg`、`poolMgr`、`customMgr`、`fetchTrigger`、`configChanged`
  - 登录密码与端口来自全局配置
- 主要导出或暴露点：
  - `New` / `Start`
  - `apiConfigSave`
  - `apiSubscriptionAdd`
  - `apiSubscriptionContribute`

## 核心流程

1. `Start()` 创建 `ServeMux`，把接口分成三层：
   - 页面入口：`/`、`/login`、`/logout`
   - 访客只读 API：如 `/api/stats`、`/api/proxies`、`/api/config`
   - 管理员 API：如 `/api/config/save`、订阅增删改刷新接口
2. `apiConfigSave()` 解析 JSON、校验关键字段、写入 `config.Save()`，再通过 `configChanged` channel 通知主流程；若池大小或协议比例变化，再触发 `poolMgr.AdjustForConfigChange()`。
3. `apiSubscriptionAdd()` 与 `apiSubscriptionContribute()` 都支持 URL 或文件内容：
   - 如有文件内容，先落到 `DATA_DIR/subscriptions`
   - 再调用 `customMgr.ValidateSubscription()` 做准入校验
   - 校验通过后入库，并异步触发 `customMgr.RefreshSubscription(id)`
4. 登录态使用包级 `sessions` map 保存，`handleLogin()` 写 cookie，`handleLogout()` 删除会话。
5. 登录页与仪表盘共享 `goproxy.theme` 主题状态：
   - 主题 key 持久化到浏览器 `localStorage`
   - 当前内置 4 套主题：`graphite-dark`、`aurora-dark`、`mist-light`、`sand-light`
   - 无保存值时按系统偏好回落到浅/深默认主题；非法值显式回退到 `graphite-dark`

## 依赖

- 数据库：SQLite，经 `storage` 读取代理、日志、订阅与配置相关数据。
- 缓存：无独立缓存；登录态保存在进程内 `sessions` map。
- 外部服务：无直接固定外部 API；订阅验证与刷新通过 `custom` 模块间接触发。
- 内部共享模块：`config`、`storage`、`pool`、`custom`、`logger`

## 测试与验证

- 已发现测试：仓库中未发现 `webui/*_test.go`。
- 已验证命令：
  - `go test ./...`
  - `go build -buildvcs=false -o .devone-runtime-data/goproxy-theme-nocgo.exe .`
- 已验证运行态：
  - 真实 `:7778` WebUI 已完成登录页 / 仪表盘主题切换、刷新持久化、guest/admin 权限显示与错误提示可读性的浏览器验证
  - 当前环境 `CGO_ENABLED=0` 时，真实冷启动通过 `storage` 包中的 `modernc.org/sqlite` `!cgo` driver 完成
- 仍未验证部分：Cookie 生命周期与长时间回归。

## 已确认 / 推断 / 未知

### 已确认

- WebUI 路由在 `Start()` 里一次性注册，访客只读接口与管理员写接口是显式分层的。（A 级证据：`webui/server.go:69-115`）
- 配置保存不会只停留在内存：`apiConfigSave()` 会调用 `config.Save()`，随后通过 `configChanged` 与 `poolMgr.AdjustForConfigChange()` 驱动主流程调整。（A 级证据：`webui/server.go:396-497`）
- 访客贡献订阅和管理员添加订阅都会先做可解析性校验，再异步刷新入池。（A 级证据：`webui/server.go:566-729`）
- 登录会话确实只保存在进程内 map，进程重启后不会持久化。（A 级证据：`webui/server.go:22-43,159-176`）
- 登录页与仪表盘现在已经共享同一套前端主题管理机制，主题状态通过 `localStorage['goproxy.theme']` 持久化，且真实页面验证已通过。（A 级证据：`webui/theme.go`、`webui/html.go`、`webui/dashboard.go` + 已接受任务包）

### 推断

- 当前 WebUI 更偏“内置管理面板”而不是独立前后端系统，因为页面模板和 API 都与主进程强耦合、一起编译发布。（B 级证据：`webui/server.go` + `webui/html.go` 组织方式）

### 未知

- 会话过期后的清理策略是否有额外后台回收逻辑，当前 `server.go` 片段中未见专门清理流程。

## 任务包同步建议

- 写入 `相关知识.md`：当前仓库没有任务包，本轮无回写目标。
- 写入 `技术说明.md`：当前仓库没有任务包，本轮无回写目标。
- 保留为升级候选：无，当前页内容已能直接升入正式 wiki。

## 来源

- `main.go:99-103`
- `webui/server.go:22-43`
- `webui/server.go:58-115`
- `webui/server.go:146-192`
- `webui/server.go:396-497`
- `webui/server.go:566-729`
- `webui/theme.go`
- `webui/html.go`
- `webui/dashboard.go`
- `.devone/data/20260405-145002-frontend-theme-system/技术说明.md`
