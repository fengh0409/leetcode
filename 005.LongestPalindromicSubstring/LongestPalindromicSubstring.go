package leetcode

// LongestPalindromicSubstring 最长回文子串
func LongestPalindromicSubstring(s string) string {
	if len(s) < 2 {
		return s
	}

	// 空间换时间，定义一个二维数组，表示字符串从l到r是否为回文串。
	// 当dp[l][r]==true时，要判断dp[l-1][r+1]是否为回文串，
	// 只需要判断字符串在l-1和r+1两个位置是否相等即可。
	// 状态转移方程：
	//   初始状态：l=r时，dp[l][r]==true
	//   状态转移：dp[l][r]==true时，且[l-1]和[r+1]两个字符相等，此时dp[l-1][r+1]==true
	var dp = make([][]bool, len(s))
	for i := range s {
		dp[i] = make([]bool, len(s))
	}
	var start, end int
	var maxLen = 1

	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > maxLen {
					maxLen = r - l + 1
					start = l
					end = r
				}
			}
		}
	}

	return s[start : end+1]
}
