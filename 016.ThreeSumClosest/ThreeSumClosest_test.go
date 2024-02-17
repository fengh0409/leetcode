package leetcode

import (
	"fmt"
	"testing"
)

type threeSumClosest struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	a      []int
	target int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one int
}

func TestThreeSumCloset(t *testing.T) {

	qs := []threeSumClosest{

		{
			input{[]int{-1, 0, 1, 1, 55}, 3},
			output{2},
		},

		{
			input{[]int{0, 0, 0}, 1},
			output{0},
		},

		{
			input{[]int{-1, 2, 1, -4}, 1},
			output{2},
		},

		{
			input{[]int{1, 1, -1}, 0},
			output{1},
		},

		{
			input{[]int{-1, 2, 1, -4}, 1},
			output{2},
		},
	}

	fmt.Printf("------------------------016.ThreeSumClosest------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, ThreeSumClosest(p.a, p.target))
	}
	fmt.Printf("\n\n\n")
}
