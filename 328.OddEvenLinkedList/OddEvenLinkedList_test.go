package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type oddEvenList struct {
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
	one []int
}

func TestOddEvenList(t *testing.T) {

	qs := []oddEvenList{

		{
			input{[]int{1, 4, 3, 2, 5, 2}},
			output{[]int{1, 3, 5, 4, 2, 2}},
		},

		{
			input{[]int{1, 1, 2, 2, 3, 3, 3}},
			output{[]int{1, 2, 3, 3, 1, 2, 3}},
		},

		{
			input{[]int{4, 3, 2, 5, 2}},
			output{[]int{4, 2, 2, 3, 5}},
		},

		{
			input{[]int{1, 1, 1, 1, 1, 1}},
			output{[]int{1, 1, 1, 1, 1, 1}},
		},

		{
			input{[]int{3, 1}},
			output{[]int{3, 1}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}},
			output{[]int{1, 3, 5, 2, 4}},
		},

		{
			input{[]int{}},
			output{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 328------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(OddEvenList(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
