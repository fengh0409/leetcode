package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type palindromeLinkedList struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one bool
}

func Test_Problem234(t *testing.T) {

	qs := []palindromeLinkedList{

		{
			input{[]int{1, 1, 2, 2, 3, 4, 4, 4}},
			output{false},
		},

		{
			input{[]int{1, 1, 1, 1, 1, 1}},
			output{true},
		},

		{
			input{[]int{1, 2, 2, 1, 3}},
			output{false},
		},

		{
			input{[]int{1}},
			output{true},
		},

		{
			input{[]int{}},
			output{true},
		},

		{
			input{[]int{1, 2, 2, 2, 2, 1}},
			output{true},
		},

		{
			input{[]int{1, 2, 2, 3, 3, 3, 3, 2, 2, 1}},
			output{true},
		},

		{
			input{[]int{1, 2}},
			output{false},
		},

		{
			input{[]int{1, 0, 1}},
			output{true},
		},

		{
			input{[]int{1, 1, 2, 1}},
			output{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 234------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, IsPalindrome(structures.Ints2List(p.one)))
	}
	fmt.Printf("\n\n\n")
}
