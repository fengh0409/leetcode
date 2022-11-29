package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type reverseLinkedList struct {
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

func TestReverseLinkedList(t *testing.T) {
	qs := []reverseLinkedList{
		{
			input{[]int{1, 2, 3, 4, 5}},
			output{[]int{5, 4, 3, 2, 1}},
		},
		{
			input{[]int{1}},
			output{[]int{1}},
		},
		{
			input{[]int{1, 2, 1, 2}},
			output{[]int{2, 1, 2, 1}},
		},
	}

	fmt.Printf("------------------------206.ReverseLinkedList------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(ReverseLinkedList(structures.Ints2List(p.one))))
	}
}
