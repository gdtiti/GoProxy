# proxy 模块

`proxy` 是 GoProxy 的对外服务层：它同时暴露 HTTP 与 SOCKS5 两套入口，按“随机轮换 / 最低延迟”模式挑选上游节点；请求失败时，免费代理直接删除，订阅代理仅禁用。

## 模块作用

- 这个模块主要负责：暴露 4 个代理端口、鉴权、上游代理选择、失败重试、HTTP 转发与 HTTPS CONNECT / SOCKS5 隧道转发。
- 当前边界：
  - `server.go`：HTTP 代理服务
  - `socks5_server.go`：SOCKS5 代理服务

## 与项目主页面的关系

- 上游页面：[架构总览](../架构总览.md)、[运行流程](../运行流程.md)
- 下游页面：[checker 模块](./checker.md)、[custom 模块](./custom.md)
- 是否属于核心路径：是，主进程最终阻塞在 `randomServer.Start()` 上。

## 入口与装配

- 启动 / 注册入口：
  - `main.go:85-91` 创建 2 个 HTTP + 2 个 SOCKS5 服务器
  - `main.go:130-154` 并发启动稳定 HTTP、稳定 SOCKS5、随机 SOCKS5，最后阻塞启动随机 HTTP（证据等级：A）
- 配置入口：
  - 构造函数接收 `cfg`
  - 运行期通过 `config.Get()` 读取 `CustomProxyMode`、优先策略、认证配置
- 主要导出或暴露点：
  - HTTP：`New` / `Start` / `ServeHTTP` / `selectProxy` / `handleHTTP` / `handleTunnel`
  - SOCKS5：`NewSOCKS5` / `Start` / `selectSOCKS5Proxy` / `socks5Handshake`

## 核心流程

1. `main.go` 为两种协议各创建“随机轮换 / 最低延迟”两个服务实例。
2. HTTP 请求进入 `ServeHTTP()` 后：
   - 如启用代理认证，先检查 `Proxy-Authorization`
   - `CONNECT` 请求走 `handleTunnel()`
   - 普通请求走 `handleHTTP()`
3. `selectProxy()` 会根据 `CustomProxyMode`、`CustomPriority`、`CustomFreePriority` 选择 `custom` / `free` / mixed 上游；失败后按 `MaxRetry` 换下一个地址重试。
4. SOCKS5 请求进入 `handleConnection()` 后，先完成 `socks5Handshake()`，再通过 `selectSOCKS5Proxy()` 只从 `protocol=socks5` 的上游集合里选代理并转发。
5. 任一服务在请求失败时都执行 `removeOrDisableProxy()`：
   - `source=custom` → `DisableProxy`
   - 其他来源 → `Delete`

## 依赖

- 数据库：SQLite，经 `storage` 读取可用代理、记录使用结果、删除/禁用故障节点。
- 缓存：无。
- 外部服务：不直接依赖固定外部 API；真正出网由上游代理完成。
- 内部共享模块：`config`、`storage`，以及 `custom` 模块注入的本地 SOCKS5 订阅节点。

## 测试与验证

- 已发现测试：仓库 `test/` 目录提供整服务级脚本，但未发现 `proxy/*_test.go`。
- 已验证命令：源码交叉验证了 HTTP / SOCKS5 两条处理链与主进程装配关系。
- 仍未验证部分：真实请求重试效果、鉴权兼容性、不同 `CustomProxyMode` 下的运行侧行为。

## 已确认 / 推断 / 未知

### 已确认

- HTTP 服务既可使用 HTTP 上游，也可使用 SOCKS5 上游；SOCKS5 服务只从 SOCKS5 上游集合里选代理。（A 级证据：`proxy/server.go:258-313`、`proxy/socks5_server.go:77-99,124-156`）
- `mixed + priority` 模式会先尝试优先来源；取不到时再 fallback 到全部来源。（A 级证据：`proxy/server.go:102-133`、`proxy/socks5_server.go:125-156`）
- 故障处理策略是“free 删除、custom 禁用”，两种协议服务共用同一处置函数。（A 级证据：`proxy/server.go:135-142` 与各自失败分支）
- 主进程最终阻塞在随机 HTTP 端口，其他三个服务通过 goroutine 先启动。（A 级证据：`main.go:130-154`）

### 推断

- `random` 端口更偏向出口多样性，`lowest-latency` 端口更偏向稳定性；这是代码与 README 共同指向的设计意图。（B 级证据：`main.go` 装配模式 + `README.md:91-116`）

### 未知

- 不同上游质量分布下，`MaxRetry` 默认值是否足够，仓库内没有运行侧统计证据。

## 任务包同步建议

- 写入 `相关知识.md`：当前仓库没有任务包，本轮无回写目标。
- 写入 `技术说明.md`：当前仓库没有任务包，本轮无回写目标。
- 保留为升级候选：无，当前页内容已能直接升入正式 wiki。

## 来源

- `main.go:85-91`
- `main.go:130-154`
- `proxy/server.go:21-64`
- `proxy/server.go:101-320`
- `proxy/socks5_server.go:15-156`
- `proxy/socks5_server.go:158-220`
- `README.md:91-116`
