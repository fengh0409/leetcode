package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type preorderTraversal struct {
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

func TestBinaryTreePreorderTraversal(t *testing.T) {

	qs := []preorderTraversal{

		{
			input{[]int{}},
			output{[]int{}},
		},

		{
			input{[]int{1}},
			output{[]int{1}},
		},

		{
			input{[]int{1, structures.NULL, 2, 3}},
			output{[]int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------144.BinaryTreePreorderTraversal------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", PreorderTraversal(root))
	}
	fmt.Printf("\n\n\n")
}
