package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type reverseKGroup struct {
	input
	output
}

// input 是参数
// one 代表第一个参数
type input struct {
	one []int
	two int
}

// output 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestReverseKGroup(t *testing.T) {

	qs := []reverseKGroup{

		{
			input{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			output{[]int{3, 2, 1, 4, 5}},
		},

		{
			input{
				[]int{1, 2, 3, 4, 5},
				1,
			},
			output{[]int{1, 2, 3, 4, 5}},
		},
	}

	fmt.Printf("------------------------025.ReverseNodesInKGroup------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(ReverseKGroup(structures.Ints2List(p.one), p.two)))
	}
	fmt.Printf("\n\n\n")
}
