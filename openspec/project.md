# Project Context

## Purpose
- 以 Go 语言实现 LeetCode 题解，用于系统化刷题与面试准备
- 强调可读性、可测试性与算法思维的沉淀

## Tech Stack
- Go 1.19
- Go Modules（`go.mod`）
- 标准库 `testing` 用于单元测试
- `vendor/github.com/halfrost/LeetCode-Go/structures` 提供通用数据结构

## Project Conventions

### Code Style
- 使用 `gofmt` 和 `go vet` 保持一致格式与静态检查
- 命名采用驼峰；包名小写不含下划线
- 目录命名 `NNN.Slug`（如 `001.TwoSum`）；文件以题名命名（如 `TwoSum.go`）
- 测试文件以 `_test.go` 结尾，表驱动测试优先

### Architecture Patterns
- 每题独立目录，包含解题代码与测试
- 通用数据结构通过 `vendor/structures` 复用（如 `ListNode`、`TreeNode`、`Heap`）
- 以纯函数实现算法，避免全局状态；必要时使用私有辅助函数

### Testing Strategy
- 运行所有测试：`go test ./... -v`
- 单题测试：`go test ./001.TwoSum -v`
- 覆盖边界输入、空输入、重复元素、极值规模等
- 对排序题提供多实现对比测试（见 `912.SortAnArray` 目录）

### Git Workflow
- 主分支：`main`；建议按题目创建分支：`feat/<id>-<slug>`
- 提交信息采用约定式：`feat: add 015 three sum`，`fix: handle empty input`
- PR 描述包含思路、复杂度与测试覆盖范围

## Domain Context
- 题型参考根文档 `README.md`：数组双指针、快速排序等
- 常用套路：双指针、二分、哈希、堆、链表操作、树的 DFS/BFS
- 复杂度目标遵循题目最优解：如 `O(n log n)` 排序、`O(n)` 线性扫描

## Important Constraints
- 仅使用标准库和 `vendor` 提供的结构；不引入外部服务
- 固定 Go 版本 `1.19`；所有代码需通过 `go test` 全量测试
- 不提交任何密钥或私密配置；遵循 `.gitignore`

## External Dependencies
- `github.com/halfrost/LeetCode-Go/structures`（经 `vendor/` 引入）：`ListNode`、`TreeNode`、`PriorityQueue`、`Heap` 等
