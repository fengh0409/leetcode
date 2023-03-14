package leetcode

import (
	"fmt"
	"testing"
)

type spiralOrder struct {
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

func TestSpiralOrder(t *testing.T) {
	qs := []spiralOrder{

		{
			input{[][]int{{3}, {2}}},
			output{[]int{3, 2}},
		},

		{
			input{[][]int{{2, 3}}},
			output{[]int{2, 3}},
		},

		{
			input{[][]int{{1}}},
			output{[]int{1}},
		},

		{
			input{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			output{[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		{
			input{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
			output{[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		},
	}

	fmt.Printf("------------------------054.SpiralMatrix------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, SpiralOrder(p.one))
	}
	fmt.Printf("\n\n\n")
}
