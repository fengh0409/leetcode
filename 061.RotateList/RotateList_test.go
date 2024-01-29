package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type rotateList struct {
	input
	output
}

// input 是参数
// one 代表第一个参数
type input struct {
	one []int
	k   int
}

// output 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestRotateRight(t *testing.T) {

	qs := []rotateList{

		{
			input{[]int{1, 2, 3, 4, 5}, 2},
			output{[]int{4, 5, 1, 2, 3}},
		},

		{
			input{[]int{1, 2, 3, 4, 5}, 3},
			output{[]int{4, 5, 1, 2, 3}},
		},

		{
			input{[]int{0, 1, 2}, 4},
			output{[]int{2, 0, 1}},
		},

		{
			input{[]int{1, 1, 1, 2}, 3},
			output{[]int{1, 1, 2, 1}},
		},

		{
			input{[]int{1}, 10},
			output{[]int{1}},
		},

		{
			input{[]int{}, 100},
			output{[]int{}},
		},
	}

	fmt.Printf("------------------------061.RotateList------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(RotateRight(structures.Ints2List(p.one), p.k)))
	}
	fmt.Printf("\n\n\n")
}
