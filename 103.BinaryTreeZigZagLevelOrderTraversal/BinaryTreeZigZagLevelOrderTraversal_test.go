package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type zigzagLevelOrder struct {
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
	one [][]int
}

func TestZigzagLevelOrder(t *testing.T) {

	qs := []zigzagLevelOrder{

		{
			input{[]int{}},
			output{[][]int{}},
		},

		{
			input{[]int{1}},
			output{[][]int{{1}}},
		},

		{
			input{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			output{[][]int{{3}, {9, 20}, {15, 7}}},
		},

		{
			input{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
			output{[][]int{{1}, {3, 2}, {4, 5}}},
		},
	}

	fmt.Printf("------------------------103.BinaryTreeZigZagLevelOrderTraversal------------------------\n")

	for _, q := range qs {
		_, p := q.output, q.input
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", ZigzagLevelOrder(root))
	}
	fmt.Printf("\n\n\n")
}
