# DevOne Worktree 台账

> 自动维护。记录每个 DevOne task worktree 的创建时间、收尾状态、收尾时间、cleanup 状态和解决的问题。
> 当最终收尾时如果 worktree 保留，必须至少更新最近收尾状态、最近收尾时间、cleanup 状态和解决的问题。

## 条目
### 20260406-232215-github-dynamic-image-publish

- 任务名称: github-dynamic-image-publish
- 任务包: .devone/data/20260406-232215-github-dynamic-image-publish
- worktree 目录: .devone/worktree/20260406-232215-github-dynamic-image-publish
- worktree 分支: devone/20260406-232215-github-dynamic-image-publish
- 开发端口: 37536
- 创建时间: 20260406-232711
- 最近收尾状态: active
- 最近收尾时间: 
- cleanup 状态: kept
- 目标分支: 
- 解决的问题: .github/workflows/docker-image.yml` 将 `IMAGE_NAME` 硬编码为 `isboyjc/goproxy`，fork 或迁移到其他账号后，GHCR 仍尝试推送到原命名空间，触发 `denied: permission_denied: The requested installation does not exist`。
- 备注: worktree 已创建，等待后续 execution / closeout

### 20260405-145002-frontend-theme-system

- 任务名称: frontend-theme-system
- 任务包: .devone/data/20260405-145002-frontend-theme-system
- worktree 目录: .devone/worktree/20260405-145002-frontend-theme-system
- worktree 分支: devone/20260405-145002-frontend-theme-system
- 开发端口: 34248
- 创建时间: 20260405-150051
- 最近收尾状态: blocked
- 最近收尾时间: 20260405-180624
- cleanup 状态: kept
- 目标分支: main
- 解决的问题: 为 GoProxy WebUI 落地统一主题系统，并兼容无 CGO 环境下的真实 SQLite 冷启动
- 备注: 已完成 completion gate 与 closeout readiness 检查；当前未执行 merge。阻塞点：主工作区 main 存在 .gitignore 修改和未跟踪 .devone/，worktree 分支也仍有未提交代码与验证产物。后续应先在 worktree 确认正式提交范围，再回到主工作区仅保留 git 集成态并重跑 worktree-merge dry-run。
