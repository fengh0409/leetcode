package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type deleteDuplicates struct {
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

func TestDeleteDuplicates(t *testing.T) {

	qs := []deleteDuplicates{

		{
			input{[]int{1, 1, 2}},
			output{[]int{1, 2}},
		},

		{
			input{[]int{1, 1, 2, 2, 3, 3, 3}},
			output{[]int{1, 2, 3}},
		},

		{
			input{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
			output{[]int{1}},
		},
	}

	fmt.Printf("------------------------083.RemoveDuplicatesFromSortedList------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(DeleteDuplicates(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
