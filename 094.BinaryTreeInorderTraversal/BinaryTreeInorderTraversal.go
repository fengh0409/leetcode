package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// InorderTraversal 二叉树的中序遍历
func InorderTraversal(root *TreeNode) (res []int) {
	// 递归终止条件
	if root == nil {
		return
	}

	// 递归访问左子树
	val := InorderTraversal(root.Left)
	res = append(res, val...)
	// 根节点
	res = append(res, root.Val)
	// 递归访问右子树
	val = InorderTraversal(root.Right)
	res = append(res, val...)
	return
}
