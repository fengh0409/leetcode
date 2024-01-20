package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type deleteDuplicates struct {
	input
	output
}

// input 是参数
// one 代表第一个参数
type input struct {
	one []int
}

// output 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestDeleteDuplicates(t *testing.T) {

	qs := []deleteDuplicates{

		{
			input{[]int{1, 1, 2, 2, 3, 4, 4, 4}},
			output{[]int{3}},
		},

		{
			input{[]int{1, 1, 1, 1, 1, 1}},
			output{[]int{}},
		},

		{
			input{[]int{1, 1, 1, 2, 3}},
			output{[]int{2, 3}},
		},

		{
			input{[]int{1}},
			output{[]int{1}},
		},

		{
			input{[]int{}},
			output{[]int{}},
		},

		{
			input{[]int{1, 2, 2, 2, 2}},
			output{[]int{1}},
		},

		{
			input{[]int{1, 1, 2, 3, 3, 4, 5, 5, 6}},
			output{[]int{2, 4, 6}},
		},

		{
			input{[]int{1, 1, 2, 3, 3, 4, 5, 6}},
			output{[]int{2, 4, 5, 6}},
		},

		{
			input{[]int{0, 1, 2, 2, 3, 4}},
			output{[]int{0, 1, 2, 2, 3, 4}},
		},
	}

	fmt.Printf("------------------------082.RemoveDuplicatesFromSortedListII------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(DeleteDuplicates(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
