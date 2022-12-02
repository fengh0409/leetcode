package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type mergeKSortedLists struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestMergeKSortedLists(t *testing.T) {
	qs := []mergeKSortedLists{

		{
			input{[][]int{}},
			output{[]int{}},
		},

		{
			input{[][]int{
				{1},
				{1},
			}},
			output{[]int{1, 1}},
		},

		{
			input{[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			}},
			output{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			input{[][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			output{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			input{[][]int{
				{1},
				{9, 9, 9, 9, 9},
			}},
			output{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			input{[][]int{
				{9, 9, 9, 9, 9},
				{1},
			}},
			output{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			input{[][]int{
				{2, 3, 4},
				{4, 5, 6},
			}},
			output{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			input{[][]int{
				{1, 3, 8},
				{1, 7},
			}},
			output{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------023.MergeKSortedLists------------------------\n")

	for _, q := range qs {
		var ls []*ListNode
		for _, qq := range q.input.one {
			ls = append(ls, structures.Ints2List(qq))
		}
		fmt.Printf("【input】:%v       【output】:%v\n", q.input.one, structures.List2Ints(MergeKLists(ls)))
	}
	fmt.Printf("\n\n\n")
}
