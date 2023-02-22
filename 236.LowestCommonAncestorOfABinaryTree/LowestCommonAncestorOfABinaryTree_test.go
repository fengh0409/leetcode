package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type lowestCommonAncestor struct {
	input
	ouput
}

// para 是参数
// one 代表第一个参数
type input struct {
	one []int
	two []int
	thr []int
}

// ans 是答案
// one 代表第一个答案
type ouput struct {
	one []int
}

func TestLowestCommonAncestor(t *testing.T) {

	qs := []lowestCommonAncestor{

		{
			input{[]int{}, []int{}, []int{}},
			ouput{[]int{}},
		},

		{
			input{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, []int{1}},
			ouput{[]int{3}},
		},

		{
			input{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, []int{4}},
			ouput{[]int{5}},
		},

		{
			input{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{8}},
			ouput{[]int{6}},
		},

		{
			input{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ouput{[]int{2}},
		},
	}

	fmt.Printf("------------------------236.LowestCommonAncestorOfABinaryTree------------------------\n")

	for _, q := range qs {
		_, p := q.ouput, q.input
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		rootTwo := structures.Ints2TreeNode(p.two)
		rootThr := structures.Ints2TreeNode(p.thr)
		fmt.Printf("【output】:%v      \n", LowestCommonAncestor(rootOne, rootTwo, rootThr))
	}
	fmt.Printf("\n\n\n")
}
