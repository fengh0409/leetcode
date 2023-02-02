package leetcode

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var rev int
	var y = x
	for y != 0 {
		n := y % 10
		rev = rev*10 + n
		y = y / 10
	}
	if rev == x {
		return true
	}
	return false
}
