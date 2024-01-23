package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type partition struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one []int
	x   int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestPartition(t *testing.T) {

	qs := []partition{

		{
			input{[]int{1, 4, 3, 2, 5, 2}, 3},
			output{[]int{1, 2, 2, 4, 3, 5}},
		},

		{
			input{[]int{1, 1, 2, 2, 3, 3, 3}, 2},
			output{[]int{1, 1, 2, 2, 3, 3, 3}},
		},

		{
			input{[]int{1, 4, 3, 2, 5, 2}, 0},
			output{[]int{1, 4, 3, 2, 5, 2}},
		},

		{
			input{[]int{4, 3, 2, 5, 2}, 3},
			output{[]int{2, 2, 4, 3, 5}},
		},

		{
			input{[]int{1, 1, 1, 1, 1, 1}, 1},
			output{[]int{1, 1, 1, 1, 1, 1}},
		},

		{
			input{[]int{3, 1}, 2},
			output{[]int{1, 3}},
		},

		{
			input{[]int{1, 2}, 3},
			output{[]int{1, 2}},
		},
	}

	fmt.Printf("------------------------086.PartitionList------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(Partition(structures.Ints2List(p.one), p.x)))
	}
	fmt.Printf("\n\n\n")
}
