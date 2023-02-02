package leetcode

// RomanToInteger 罗马数字转整数
func RomanToInteger(s string) int {
	var sum int
	var m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// 把一个小值放在大值的左边，就是做减法，否则为加法。
	// 比较当前值与后一位的值，当前值比后一位小小于则做减法，大于则做加法
	for i := 0; i < len(s)-1; i++ {
		if m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}
	}
	// 最后一个元素做加法
	sum += m[s[len(s)-1]]

	return sum

	//var m = map[byte]int{
	//	'I': 1,
	//	'V': 5,
	//	'X': 10,
	//	'L': 50,
	//	'C': 100,
	//	'D': 500,
	//	'M': 1000,
	//}
	//var n = map[string]int{
	//	"IV": 4,
	//	"IX": 9,
	//	"XL": 40,
	//	"XC": 90,
	//	"CD": 400,
	//	"CM": 900,
	//}

	//var sum int
	//for i := 0; i < len(s); i++ {
	//	if i+1 < len(s) {
	//		if v, ok := n[string(s[i])+string(s[i+1])]; ok {
	//			sum += v
	//			i++
	//		} else {
	//			sum += m[s[i]]
	//		}
	//	} else {
	//		sum += m[s[i]]
	//	}
	//}

	//return sum
}
