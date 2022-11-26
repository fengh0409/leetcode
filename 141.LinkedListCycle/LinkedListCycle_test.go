package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type linkedListCycle struct {
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
	one bool
}

func TestHasCycle(t *testing.T) {
	qs := []linkedListCycle{

		{
			input{[]int{3, 2, 0, -4}},
			output{false},
		},

		{
			input{[]int{1, 2}},
			output{false},
		},

		{
			input{[]int{1}},
			output{false},
		},
	}

	fmt.Printf("------------------------141.LinkedListCycle------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, HasCycle(structures.Ints2List(p.one)))
	}
	fmt.Printf("\n\n\n")
}
