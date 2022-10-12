package leetcode

var combinations []string
var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// LetterCombinationsOfAPhoneNumber 电话号码的字母组合
func LetterCombinationsOfAPhoneNumber(digits string) []string {
	if digits == "" {
		return []string{}
	}

	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

// 回溯法
func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		number := string(digits[index])
		letter := m[number]
		for i := 0; i < len(letter); i++ {
			backtrack(digits, index+1, combination+string(letter[i]))
		}
	}
}
