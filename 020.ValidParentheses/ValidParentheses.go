package leetcode

// ValidParentheses 有效的括号
func ValidParentheses(s string) bool {
	var m = map[byte]byte{'(': ')', '{': '}', '[': ']'}
	// 遇到左括号入栈，遇到右括号时判断栈顶元素是否为对应的左括号，是则把该栈顶元素出栈，遍历结束后，栈为空，则视为合法的括号。
	var stack = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			// 若为左括号入栈
			stack = append(stack, s[i])
		} else {
			// 若为右括号，则移出栈顶元素
			if len(stack) > 0 && s[i] == m[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				// 如果上面都if没通过，则肯定是非法都括号，直接return false
				return false
			}
		}
	}

	return len(stack) == 0
}
