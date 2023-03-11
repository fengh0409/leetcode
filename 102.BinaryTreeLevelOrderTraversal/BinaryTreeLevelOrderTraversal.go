package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// LevelOrder 二叉树的层序遍历，从上到下，从左到右遍历
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var tmpVal []int
		var tmpQueue = queue
		// 这里需要将queue清空，重新赋值
		queue = []*TreeNode{}
		for _, node := range tmpQueue {
			tmpVal = append(tmpVal, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmpVal)
	}

	return result
}
