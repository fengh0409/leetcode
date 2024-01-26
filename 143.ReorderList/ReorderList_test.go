package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type reorderList struct {
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

func TestReorderList(t *testing.T) {

	qs := []reorderList{

		{
			input{[]int{1, 2, 3, 4, 5}},
			output{[]int{1, 5, 2, 4, 3}},
		},
		{
			input{[]int{1, 2, 3, 4}},
			output{[]int{1, 4, 2, 3}},
		},

		{
			input{[]int{1}},
			output{[]int{1}},
		},

		{
			input{[]int{}},
			output{[]int{}},
		},
	}

	fmt.Printf("------------------------143.ReorderList-----------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		ReorderList(structures.Ints2List(p.one))
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(structures.Ints2List(p.one)))
	}
	fmt.Printf("\n\n\n")
}
