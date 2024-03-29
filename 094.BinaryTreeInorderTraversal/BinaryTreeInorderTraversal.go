package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// InorderTraversal 二叉树的中序遍历
// 二叉树的前序、中序、后序遍历是以根节点来区分的，前序遍历即先读取根节点，再遍历左节点，最后遍历右节点
// 中序遍历即根节点放在中间遍历，先遍历左节点，再到根节点，最后遍历右节点
// 后序遍历即最后遍历根节点，先遍历左节点，再遍历右节点，最后到根节点
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
