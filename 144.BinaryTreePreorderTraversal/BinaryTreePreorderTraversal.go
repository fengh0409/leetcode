package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// InorderTraversal 二叉树的前序遍历
// 二叉树的前序、中序、后序遍历是以根节点来区分的，前序遍历即先读取根节点，再遍历左节点，最后遍历右节点
// 中序遍历即根节点放在中间遍历，先遍历左节点，再到根节点，最后遍历右节点
// 后序遍历即最后遍历根节点，先遍历左节点，再遍历右节点，最后到根节点
func PreorderTraversal(root *TreeNode) []int {
	var result = []int{}
	traversal(root, &result)
	return result

	// 迭代法，显示声明一个栈，栈是先入后出，前序遍历是先左节点再右节点，则是右节点先入栈，左节点后入栈
	// if root == nil {
	// 	return []int{}
	// }
	// var res []int
	// var stack = []*TreeNode{}
	// stack = append(stack, root)
	// for len(stack) != 0 {
	// 	node := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	if node != nil {
	// 		res = append(res, node.Val)
	// 	}
	// 	if node.Right != nil {
	// 		stack = append(stack, node.Right)
	// 	}
	// 	if node.Left != nil {
	// 		stack = append(stack, node.Left)
	// 	}
	// }
	// return res
}

func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	traversal(root.Left, result)
	traversal(root.Right, result)
}
