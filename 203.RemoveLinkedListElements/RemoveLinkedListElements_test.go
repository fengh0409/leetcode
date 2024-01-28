package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type removeElements struct {
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

func TestRemoveElements(t *testing.T) {

	qs := []removeElements{

		{
			input{[]int{1, 2, 3, 4, 5}, 1},
			output{[]int{2, 3, 4, 5}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 2},
			output{[]int{1, 3, 4, 5}},
		},

		{
			input{[]int{1, 1, 1, 1, 1}, 1},
			output{[]int{}},
		},

		{
			input{[]int{1, 2, 3, 2, 3, 2, 3, 2}, 2},
			output{[]int{1, 3, 3, 3}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 5},
			output{[]int{1, 2, 3, 4}},
		},

		{
			input{[]int{}, 5},
			output{[]int{}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 10},
			output{[]int{1, 2, 3, 4, 5}},
		},

		{
			input{[]int{1}, 1},
			output{[]int{}},
		},
	}

	fmt.Printf("------------------------203.RemoveLinkedListElements------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(RemoveElements(structures.Ints2List(p.one), p.n)))
	}
	fmt.Printf("\n\n\n")
}
