package leetcode

// SpiralOrder 螺旋矩阵，从左往右，从上往下，从右往左，从下往上，依次输出元素
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	up, down := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := make([]int, 0)
	for {
		// 从左往右遍历
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}
		// 最上边一层遍历完成，那么上边界up要加一
		up++
		if up > down {
			break
		}

		// 从上往下遍历
		for i := up; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		// 最右边一列遍历完成，那么右边界right要减一
		right--
		if right < left {
			break
		}

		// 从右往左遍历
		for i := right; i >= left; i-- {
			result = append(result, matrix[down][i])
		}
		// 最下边一层遍历完成，那么下边界down要减一
		down--
		if down < up {
			break
		}

		// 从下往上遍历
		for i := down; i >= up; i-- {
			result = append(result, matrix[i][left])
		}
		// 最左边一列遍历完成，那么左边界left要加一
		left++
		if left > right {
			break
		}
	}

	return result
}
