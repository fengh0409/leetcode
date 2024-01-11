package leetcode

// LengthOfLongestSubstring 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	// 滑动窗口解法，时间复杂度O(n)，从第一个字符开始，窗口右边界往后移动（可以理解为指针），
	// 如果遇到重复字符，则把窗口左边界缩小到前面的重复字符之后（相当于把前面重复字符及其左边的字符全部移出了窗口）
	// 例如：字符 abcabdeccce，窗口右边界从第一个字符 a 开始往后移动，到第二个字符 a 之前（此时窗口内容为abc），窗口大小为 3，
	// 即最长子串为 3，遇到第二个字符 a 后，窗口左边界移到第一个字符 a 的后面（即第一个字符 a 被移出了窗口，此时窗口内容为bca），
	// 此时最长子串也为 3，窗口右边界继续向后移动，每次移动都更新最长子串，直至字符串末尾。
	// 在程序实现时，子串长度的计算方式为：右边界的下标-左边界的下标+1 = 子串的长度
	// 定义左右边界
	var left, right, maxSubstrLen int
	var m = map[byte]int{}
	for k, v := range []byte(s) {
		if k0, ok := m[v]; ok {
			// 若遇到重复字符，则将滑动窗口左边界向右移动，从而缩小窗口
			// !!!注意：这里一定要判断 k0 >= left
			// 考虑字符串bcaabd，当遍历到第二个a时，由于a重复了，所以滑动窗口缩小窗口，bca被移出了滑动窗口，
			// 此时仅包含字符a，窗口的左右边界都是第二个a的下标，即left=right=3，此时m=map[byte]int{'b':0,'c':1,'a':3}。
			// 当遍历到第二个b时，若没有k0>=left的判断，则left=0+1=1，窗口的左边界往左边移动了！！
			// 此时子串长度为right-left+1=4-1+1=4，这是错误的，因为bca都不在窗口中了。
			// 所以当遇到重复字符时，要判断之前的字符位置k0与left的大小，如果k0<left，说明之前的字符已经在窗口之外了，
			// 这时不能更改left，仅当k0>=left时，左边界才向右移动一位。
			if k0 >= left {
				// 缩小左边窗口，向右边移动一位
				left = k0 + 1
			}
		}
		// 每次遍历，右边界即为当前下标
		right = k
		if right-left+1 > maxSubstrLen {
			maxSubstrLen = right - left + 1
		}
		// 将字符作为key和下标作为value存到map
		m[v] = k
	}

	return maxSubstrLen

	// 两次遍历，时间复杂度O(n²)
	//var maxSubstrLen int
	//parts := strings.Split(s, "")
	//for i := 0; i < len(parts); i++ {
	//	var substr = parts[i]
	//	for j := i + 1; j < len(parts); j++ {
	//		if strings.Contains(substr, parts[j]) {
	//			break
	//		}
	//		substr += parts[j]
	//	}

	//	if len(substr) > maxSubstrLen {
	//		maxSubstrLen = len(substr)
	//	}
	//}

	//return maxSubstrLen
}
