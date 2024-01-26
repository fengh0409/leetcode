package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type middleNode struct {
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
	one int
}

func TestMiddleNode(t *testing.T) {

	qs := []middleNode{

		{
			input{[]int{1, 2, 3, 4, 5}},
			output{3},
		},
		{
			input{[]int{1, 2, 3, 4}},
			output{3},
		},

		{
			input{[]int{1}},
			output{1},
		},

		{
			input{[]int{}},
			output{},
		},
	}

	fmt.Printf("------------------------876.MiddleNode------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(MiddleNode(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
