package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type mergeTwoSortedLists struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestMergeTwoSortedLists(t *testing.T) {
	qs := []mergeTwoSortedLists{
		{
			input{[]int{}, []int{}},
			output{[]int{}},
		},

		{
			input{[]int{1}, []int{1}},
			output{[]int{1, 1}},
		},

		{
			input{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			output{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			output{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			input{[]int{1}, []int{9, 9, 9, 9, 9}},
			output{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			input{[]int{9, 9, 9, 9, 9}, []int{1}},
			output{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			input{[]int{2, 3, 4}, []int{4, 5, 6}},
			output{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			input{[]int{1, 3, 8}, []int{1, 7}},
			output{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------021.MergeTwoSortedLists------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(MergeTwoSortedLists(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
