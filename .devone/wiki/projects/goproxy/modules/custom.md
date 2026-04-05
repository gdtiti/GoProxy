# custom 模块

`custom` 负责“订阅进来之后怎样变成可用代理”：直连 HTTP/SOCKS5 节点直接入池；需要隧道转换的加密节点统一交给 `sing-box` 子进程转成 `127.0.0.1:{port}` 的本地 SOCKS5 代理，再回写到代理池。

## 模块作用

- 这个模块主要负责：订阅拉取、节点解析、节点分流、`sing-box` 启停/重载、订阅代理验证、定时刷新与探测唤醒。
- 当前边界：
  - `manager.go`：订阅生命周期与后台循环。
  - `parser.go`：订阅格式解析。
  - `singbox.go`：`sing-box` 进程与端口映射管理。

## 与项目主页面的关系

- 上游页面：[架构总览](../架构总览.md)、[运行流程](../运行流程.md)
- 下游页面：[proxy 模块](./proxy.md)、[webui 模块](./webui.md)、[开放问题](../开放问题.md)
- 是否属于核心路径：是，既参与启动期后台循环，也被 WebUI 订阅接口直接调用。

## 入口与装配

- 启动 / 注册入口：
  - `main.go:93-95` 创建 `custom.NewManager(...)`
  - `main.go:124-125` 启动 `go customMgr.Start()`
  - `webui/server.go:102-105,642-724` 由 WebUI 订阅接口触发 `ValidateSubscription()` / `RefreshSubscription()`。（证据等级：A）
- 配置入口：
  - `NewManager()` 读取 `DATA_DIR`、`cfg.SingBoxPath`、`cfg.SingBoxBasePort`
  - 刷新与探测间隔来自全局配置
- 主要导出或暴露点：
  - `NewManager` / `Start` / `RefreshSubscription`
  - `ValidateSubscription` / `GetStatus`
  - `SingBoxProcess.startLocked` / `GetPortMap` / `GetNodeCount`

## 核心流程

1. `Start()` 启动三个后台分支：`initialRefresh()`、`refreshLoop()`、`probeLoop()`。
2. `RefreshSubscription(subID)` 读取订阅、拉取内容、调用 `Parse()` 解析节点，并先清掉该订阅旧代理。
3. 节点被拆成两类：
   - `IsDirect()` 为真：直接 `AddProxyWithSource(..., "custom", subID)` 入池
   - 其余节点：合并所有 tunnel 节点后交给 `singbox.Reload()`，再通过 `GetPortMap()` 取回本地端口映射，写成 `127.0.0.1:{port}` 的 SOCKS5 代理入池
4. 新代理随后进入 `validateCustomProxies()`：
   - 验证成功且未触发地理过滤 → `EnableProxy`
   - 验证失败或地理受限 → `DisableProxy`
   - 只要存在可用节点 → `UpdateSubscriptionSuccess(subID)`

## 依赖

- 数据库：SQLite，经 `storage` 管理订阅、代理、启用/禁用状态。
- 缓存：无独立缓存。
- 外部服务：
  - 订阅 URL 拉取
  - `sing-box` 子进程与本地监听端口
- 内部共享模块：`config`、`storage`、`validator`

## 测试与验证

- 已发现测试：仓库中未发现 `custom/*_test.go`。
- 已验证命令：源码层交叉验证了 `main.go`、`webui/server.go`、`custom/manager.go`、`custom/singbox.go` 的调用链。
- 仍未验证部分：真实订阅下载、`sing-box` 启动兼容性、探测唤醒与 7 天无可用节点清理的运行侧效果。

## 已确认 / 推断 / 未知

### 已确认

- `custom.Manager` 是主流程唯一装配的订阅管理入口。（A 级证据：`main.go:93-95,124-125`）
- 直连节点直接入池；加密节点经 `sing-box` 重载后映射为本地 SOCKS5 端口再入池。（A 级证据：`custom/manager.go:245-307`、`custom/singbox.go:362-443,491-507`）
- `ValidateSubscription()` 只负责“能否解析出节点”的准入校验，不直接入库。（A 级证据：`custom/manager.go:511-539`）
- WebUI 的管理员添加订阅和访客贡献订阅，都会先做 `ValidateSubscription()`，再异步触发 `RefreshSubscription()`。（A 级证据：`webui/server.go:606-651,698-729`）

### 推断

- 当前 `sing-box` 不是“每个订阅一个进程”，而是把全部 tunnel 节点合并后交给同一进程统一管理。（B 级证据：`RefreshSubscription()` 中 `collectAllTunnelNodes()` + `Reload(mergedNodes)`）

### 未知

- `Stop()` 在当前主流程里没有显式接入优雅退出，仓库内也没有统一 shutdown 编排说明。

## 任务包同步建议

- 写入 `相关知识.md`：当前仓库没有任务包，本轮无回写目标。
- 写入 `技术说明.md`：当前仓库没有任务包，本轮无回写目标。
- 保留为升级候选：无，当前页内容已能直接升入正式 wiki。

## 来源

- `main.go:93-95`
- `main.go:124-125`
- `custom/manager.go:22-58`
- `custom/manager.go:211-313`
- `custom/manager.go:458-539`
- `custom/singbox.go:362-443`
- `custom/singbox.go:491-507`
- `webui/server.go:99-105`
- `webui/server.go:566-729`
