package leetcode

// ValidParentheses 有效的括号
func ValidParentheses(s string) bool {
	// 遇到左括号入栈，遇到右括号判断栈顶元素是否为对应的左括号，是则把该栈顶元素出栈，遍历结束后，栈为空，则视为合法的括号。
	if len(s)%2 == 1 {
		return false
	}

	var stack = make([]byte, 0)
	// 左右括号分别以key、value存储在map中
	var m = map[byte]byte{'(': ')', '{': '}', '[': ']'}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			// 左括号入栈
			stack = append(stack, s[i])
		} else {
			// 右括号出栈
			if len(stack) > 0 && s[i] == m[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
