package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type swapPairs struct {
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

func TestSwapPairs(t *testing.T) {

	qs := []swapPairs{

		{
			input{[]int{}},
			output{[]int{}},
		},

		{
			input{[]int{1}},
			output{[]int{1}},
		},

		{
			input{[]int{1, 2, 3, 4}},
			output{[]int{2, 1, 4, 3}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}},
			output{[]int{2, 1, 4, 3, 5}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------024.SwapNodesInPairs------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(SwapPairs(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
