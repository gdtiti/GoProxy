# 模块索引

GoProxy 的主路径模块集中在 `checker`、`custom`、`proxy`、`webui` 四块：前两者负责后台维护与订阅接入，后两者负责对外服务与管理入口。这里不重复项目总览，只补组件级稳定知识。

## 阅读建议

1. 先读 [proxy 模块](./proxy.md)，理解对外 HTTP / SOCKS5 服务怎样选上游、怎样失败重试。
2. 再读 [custom 模块](./custom.md)，理解订阅节点怎样入池，以及 `sing-box` 在哪里接入。
3. 接着读 [checker 模块](./checker.md)，理解 free 删除 / custom 禁用的维护策略。
4. 最后读 [webui 模块](./webui.md)，理解访客只读、管理员写操作和订阅入口。

## 模块页导航

- [checker 模块](./checker.md)
- [custom 模块](./custom.md)
- [proxy 模块](./proxy.md)
- [webui 模块](./webui.md)

## 分层说明

- 这里只收录已经能从当前仓库主线直接验证的模块边界与调用链。
- 仍依赖外部发布平台、其他分支或运行环境才能确认的内容，保留在项目级 [开放问题](../开放问题.md)。

## 已确认 / 推断 / 未知

### 已确认

- `checker`、`custom`、`proxy`、`webui` 四页都能直接回溯到 `main.go` 的装配链。（A 级证据：`main.go:53-154`）
- 这四块覆盖了当前仓库最关键的“后台维护 → 订阅接入 → 对外代理 → 管理入口”主路径。（A 级证据：`README.md:240-255` + 各模块源码）

### 推断

- 对多数读者而言，先读 `proxy` 和 `custom`，再读 `checker` 与 `webui`，更容易建立系统整体心智模型。（B 级证据：基于当前调用链与职责依赖的阅读建议）

### 未知

- 其余目录如 `fetcher`、`pool`、`storage` 是否需要继续拆成独立模块页，要看后续是否出现更细的阅读/交接需求。

## 来源

- `main.go:53-154`
- `README.md:240-255`
- `checker/health_checker.go:13-108`
- `custom/manager.go:22-58`
- `proxy/server.go:21-320`
- `webui/server.go:58-115`
