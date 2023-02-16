package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type levelOrder struct {
	input
	ouput
}

// one 代表第一个参数
type input struct {
	one []int
}

// one 代表第一个答案
type ouput struct {
	one [][]int
}

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {

	qs := []levelOrder{

		{
			input{[]int{}},
			ouput{[][]int{}},
		},

		{
			input{[]int{1}},
			ouput{[][]int{{1}}},
		},

		{
			input{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ouput{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	fmt.Printf("------------------------102.BinaryTreeLevelOrderTraversal------------------------\n")

	for _, q := range qs {
		_, p := q.ouput, q.input
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", LevelOrder(root))
	}
	fmt.Printf("\n\n\n")
}
