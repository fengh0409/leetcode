package leetcode

// IsPalindrome 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// 先反转整数
	var rev int
	var y = x
	for y != 0 {
		n := y % 10
		rev = rev*10 + n
		y = y / 10
	}
	// 再比较数据是否相等
	if rev == x {
		return true
	}
	return false
}
