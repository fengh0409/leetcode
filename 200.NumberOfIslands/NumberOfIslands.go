package leetcode

// NumIslands 岛屿数量
// 参考题解：https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
func NumIslands(grid [][]byte) int {
	var nums = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				nums++
			}
		}
	}
	return nums
}

// dfs 深度优先搜索
func dfs(grid [][]byte, i, j int) {
	// 如果遍历的元素不在范围内，则表示越界了，直接返回
	if !inArea(grid, i, j) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	// 遍历后的元素做个标记，以免被重复遍历
	grid[i][j] = 2
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// inArea 判断某个元素是否在二维数组内
func inArea(grid [][]byte, i, j int) bool {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		return true
	}
	return false
}
