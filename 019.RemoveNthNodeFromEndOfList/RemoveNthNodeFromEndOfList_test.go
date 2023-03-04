package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type removeNthFromEnd struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestRemoveNthFromEnd(t *testing.T) {

	qs := []removeNthFromEnd{

		{
			input{[]int{1}, 3},
			output{[]int{1}},
		},

		{
			input{[]int{1, 2}, 2},
			output{[]int{2}},
		},

		{
			input{[]int{1}, 1},
			output{[]int{}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 10},
			output{[]int{1, 2, 3, 4, 5}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 1},
			output{[]int{1, 2, 3, 4}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 2},
			output{[]int{1, 2, 3, 5}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 3},
			output{[]int{1, 2, 4, 5}},
		},
		{
			input{[]int{1, 2, 3, 4, 5}, 4},
			output{[]int{1, 3, 4, 5}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 5},
			output{[]int{2, 3, 4, 5}},
		},
	}

	fmt.Printf("------------------------019.RemoveNthNodeFromEndOfList------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(RemoveNthFromEnd(structures.Ints2List(p.one), p.n)))
	}
	fmt.Printf("\n\n\n")
}