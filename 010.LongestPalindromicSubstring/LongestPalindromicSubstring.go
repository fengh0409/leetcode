package leetcode

// LongestPalindromicSubstring 最长回文子串
func LongestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	var mid = (len(s) + 1) / 2
	var substr string
	var start, end int
	var restart bool
	for i := 0; i < len(s); i++ {
		if restart {
			start = i
			end = len(s) - i - 1
		}

		for j := len(s) - i - 1; j >= i; j-- {
			if s[i] == s[j] {
				if j-i == 1 {
					restart = true
				}
				if j == i {
					restart = true
				}
				break
			}

		}

		// 遍历到中间位
		if restart {
			substr = s[start : end+1]
		}
	}

	return substr
}
