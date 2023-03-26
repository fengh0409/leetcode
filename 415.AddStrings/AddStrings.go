package leetcode

import (
	"strconv"
)

// AddStrings 字符串相加
func AddStrings(num1 string, num2 string) string {
	var result string
	var add int
	i, j := len(num1)-1, len(num2)-1
	for ; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			// 这里 - ‘0’ 是ascii码转整型
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + add
		// 注意这里是字符串拼接
		result = strconv.Itoa(sum%10) + result
		add = sum / 10
	}
	return result
}
