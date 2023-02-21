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
		tmpQueue := queue
		queue = nil
		length := len(tmpQueue)
		tmp := make([]int, 0, length)
		for i := 0; i < length; i++ {
			tmp = append(tmp, tmpQueue[i].Val)
			if tmpQueue[i].Left != nil {
				queue = append(queue, tmpQueue[i].Left)
			}
			if tmpQueue[i].Right != nil {
				queue = append(queue, tmpQueue[i].Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			n := len(tmp)
			for i := 0; i < n/2; i++ {
				tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
			}
		}
		result = append(result, tmp)
	}
	return result
}
