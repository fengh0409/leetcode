package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// ZigzagLevelOrder 二叉树的锯齿形层序遍历，从上到下，从左往右遍历，再从右往左遍历
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var tmpVal []int
		var tmpQueue = queue
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
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			n := len(tmpVal)
			for i := 0; i < n/2; i++ {
				tmpVal[i], tmpVal[n-1-i] = tmpVal[n-1-i], tmpVal[i]
			}
		}
		result = append(result, tmpVal)
	}
	return result
}
