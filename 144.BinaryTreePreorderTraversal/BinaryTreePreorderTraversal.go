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
}

func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	traversal(root.Left, result)
	traversal(root.Right, result)
}
