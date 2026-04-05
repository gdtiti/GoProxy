# checker 模块

`checker` 是 GoProxy 的后台健康维护模块：当前主链真正运行的是 `HealthChecker`，它按批次复查代理并执行 `free 删除 / custom 禁用`；同目录里还保留了一套旧 `Checker` 实现，但在当前 `main.go` 装配链中没有被使用。

## 模块作用

- 这个模块主要负责：周期性健康检查、延迟/出口信息回写、失败计数与失效节点处置。
- 当前边界：
  - `checker/health_checker.go`：当前主流程使用的实现。
  - `checker/checker.go`：仓库内仍保留的旧实现。

## 与项目主页面的关系

- 上游页面：[架构总览](../架构总览.md)、[运行流程](../运行流程.md)
- 下游页面：[开放问题](../开放问题.md)
- 是否属于核心路径：是，直接挂在 `main.go` 后台任务链上。

## 入口与装配

- 启动 / 注册入口：`main.go:58-59` 创建 `checker.NewHealthChecker(...)`，`main.go:118-119` 调用 `healthChecker.StartBackground()`。（证据等级：A）
- 配置入口：`HealthChecker` 在构造时接收 `cfg`，旧 `Checker` 在循环内调用 `config.Get()` 读取最新配置。
- 主要导出或暴露点：
  - `NewHealthChecker` / `RunOnce` / `StartBackground`
  - 旧实现 `New` / `Start`

## 核心流程

1. `StartBackground()` 按 `HealthCheckInterval` 创建 ticker，定时调用 `RunOnce()`。
2. `RunOnce()` 先通过 `poolMgr.GetStatus()` 与 `storage.GetQualityDistribution()` 判断是否跳过部分 S 级代理。
3. `storage.GetBatchForHealthCheck()` 拉取待检查批次，`validator.ValidateStream()` 并发验证。
4. 成功结果通过 `storage.UpdateExitInfo()` 回写出口信息与延迟；失败结果先 `IncrementFailCount()`，达到阈值后：
   - `source=custom` → `DisableProxy`
   - 其他来源 → `Delete`

## 依赖

- 数据库：SQLite，经 `storage.Storage` 访问代理记录与失败计数。
- 缓存：无。
- 外部服务：不直接调用；外部探测由 `validator` 间接完成。
- 内部共享模块：`config`、`pool`、`storage`、`validator`。

## 测试与验证

- 已发现测试：仓库中未发现 `checker/*_test.go`。
- 已验证命令：源码交叉验证了 `main.go` 装配链和 `checker` 两套实现的入口关系。
- 仍未验证部分：运行时健康检查节奏、批次选择效果、失败阈值触发后的真实清理结果。

## 已确认 / 推断 / 未知

### 已确认

- `HealthChecker` 才是当前 `main.go` 主链启动的实现。（A 级证据：`main.go:58-59,118-119`）
- 当前健康检查策略不是“一次全量扫库”，而是“按批次抽取 + 流式验证 + 达阈值再移除/禁用”。（A 级证据：`checker/health_checker.go:30-108`）
- 旧 `Checker` 的 `New()` / `Start()` 在当前 `main.go` 里没有被引用。（A 级证据：`main.go:53-154` 对照 `checker/checker.go:16-29`）

### 推断

- `checker/checker.go` 更像迁移过程中留下的旧实现痕迹，而不是当前主链必需组件。（B 级证据：当前仓库主装配链未引用，但未见作者说明）

### 未知

- 旧 `Checker` 是否仍被仓库外脚本、其他分支或人工调用依赖，当前仓库内没有证据。

## 任务包同步建议

- 写入 `相关知识.md`：当前仓库没有任务包，本轮无回写目标。
- 写入 `技术说明.md`：当前仓库没有任务包，本轮无回写目标。
- 保留为升级候选：无，当前页内容已能直接升入正式 wiki。

## 来源

- `main.go:53-59`
- `main.go:118-119`
- `checker/checker.go:12-29`
- `checker/checker.go:31-78`
- `checker/health_checker.go:13-108`
