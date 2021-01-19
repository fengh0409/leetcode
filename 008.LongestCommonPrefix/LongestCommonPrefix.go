package leetcode

// LongestCommonPrefix 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// 时间复杂度O(mn)
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var prefix = strs[0]
	// 将第一个字符串作为最长前缀，从第二个元素开始遍历数组，将其与prefix进行比较，得到新的prefix
	for i := 1; i < len(strs); i++ {
		var length = len(prefix)
		if len(prefix) > len(strs[i]) {
			length = len(strs[i])
		}
		j := 0
		for ; j < length; j++ {
			if prefix[j] != strs[i][j] {
				break
			}
		}
		// 这里字符串拼接比较耗时，更好的做法是 prefix = prefix[:j]
		//var temp string
		//for j := 0; j < length; j++ {
		//	if prefix[j] != strs[i][j] {
		//		break
		//	}
		//	temp += string(prefix[j])
		//}
		//prefix = temp
		prefix = prefix[:j]
		if prefix == "" {
			break
		}
	}

	return prefix
}
