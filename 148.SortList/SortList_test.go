package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type sortList struct {
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

func TestSortList(t *testing.T) {
	qs := []sortList{
		{
			input{[]int{1, 2, 3, 4, 5}},
			output{[]int{1, 2, 3, 4, 5}},
		},
		{
			input{[]int{1, 1, 2, 5, 5, 4, 10, 0}},
			output{[]int{0, 1, 1, 2, 4, 5, 5, 10}},
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

	fmt.Printf("------------------------148.SortList------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(SortList(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
