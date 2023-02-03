package leetcode

// LongestCommonPrefix 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 纵向扫描，依次遍历每个字符串，更新最长公共前缀
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			// 这里len(strs[i]) == j是避免strs[i][j]取值时越界
			if len(strs[i]) == j || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
