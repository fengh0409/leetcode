package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type intersectionOfTwoLinkedLists struct {
	input
	output
}

// para 是参数
// one 代表第一个参数
type input struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type output struct {
	one []int
}

func TestGetIntersectionNode(t *testing.T) {

	qs := []intersectionOfTwoLinkedLists{

		{
			input{[]int{}, []int{}},
			output{[]int{}},
		},

		{
			input{[]int{3}, []int{1, 2, 3}},
			output{[]int{3}},
		},

		{
			input{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			output{[]int{1, 2, 3, 4}},
		},

		{
			input{[]int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}},
			output{[]int{8, 4, 5}},
		},

		{
			input{[]int{1}, []int{9, 9, 9, 9, 9}},
			output{[]int{}},
		},

		{
			input{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}},
			output{[]int{2, 4}},
		},
	}

	fmt.Printf("------------------------160.IntersectionOfTwoLinkedLists------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(GetIntersectionNode(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
