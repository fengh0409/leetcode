package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type reverseBetween struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one  []int
	m, n int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestReverseLinkedListII(t *testing.T) {

	qs := []reverseBetween{

		{
			input{[]int{1, 2, 3, 4, 5}, 2, 4},
			output{[]int{1, 4, 3, 2, 5}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 2, 2},
			output{[]int{1, 2, 3, 4, 5}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 1, 5},
			output{[]int{5, 4, 3, 2, 1}},
		},

		{
			input{[]int{1, 2, 3, 4, 5, 6}, 3, 4},
			output{[]int{1, 2, 4, 3, 5, 6}},
		},

		{
			input{[]int{3, 5}, 1, 2},
			output{[]int{5, 3}},
		},
	}

	fmt.Printf("------------------------092.ReverseLinkedListII------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(ReverseBetween(structures.Ints2List(p.one), p.m, p.n)))
	}
	fmt.Printf("\n\n\n")
}
