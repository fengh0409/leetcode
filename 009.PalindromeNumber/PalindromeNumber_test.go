package leetcode

import (
	"fmt"
	"testing"
)

type isPalindrome struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one bool
}

func TestIsPalindrome(t *testing.T) {

	qs := []isPalindrome{

		{
			input{121},
			output{true},
		},

		{
			input{-121},
			output{false},
		},

		{
			input{10},
			output{false},
		},

		{
			input{321},
			output{false},
		},

		{
			input{-123},
			output{false},
		},

		{
			input{120},
			output{false},
		},

		{
			input{1534236469},
			output{false},
		},
	}

	fmt.Printf("------------------------009.PalindromeNumber------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, IsPalindrome(p.one))
	}
	fmt.Printf("\n\n\n")
}
