# Change: 添加数组原地移动零能力

## Why
LeetCode 常见需求（283 Move Zeroes）：在不复制数组的前提下，将所有 0 移动至末尾并保持非零元素的相对顺序。该能力可沉淀为统一的数组处理规范，便于后续题解与复用。

## What Changes
- 新增能力规范：`arrays/move-zeroes`（原地、稳定、线性时间）
- 约定函数签名：`func MoveZeroes(nums []int)`，原地修改 `nums`
- 明确复杂度：时间 `O(n)`、额外空间 `O(1)`、保持非零元素相对顺序
- 明确测试覆盖：空数组、全零、无零、交错零与负数、典型示例
- 计划在实现阶段新增目录 `283.MoveZeroes/` 与配套测试

## Impact
- Affected specs: `arrays/move-zeroes`
- Affected code: 新增 `283.MoveZeroes/MoveZeroes.go`、`283.MoveZeroes/MoveZeroes_test.go`

