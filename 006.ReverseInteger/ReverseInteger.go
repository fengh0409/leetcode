package leetcode

// ReverseInteger 反转整数
func ReverseInteger(x int) int {
	// 时间复杂度 O(log10(x))
	var rev int
	for x != 0 {
		residual := x % 10
		rev = rev*10 + residual
		if rev > 1<<31-1 || rev < -1<<31 {
			return 0
		}
		x = x / 10
	}

	return rev
}
