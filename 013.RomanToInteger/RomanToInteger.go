package leetcode

// RomanToInteger 罗马数字转整数
func RomanToInteger(s string) int {
	// 注意总结规律
	if len(s) == 0 {
		return 0
	}
	var m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum int
	var pre = m[s[0]]
	for i := 1; i < len(s); i++ {
		if pre < m[s[i]] {
			sum -= pre
		} else {
			sum += pre
		}
		pre = m[s[i]]
	}
	// 加最后一个元素
	sum += pre

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
