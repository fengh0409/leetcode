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
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, 0, length)
		for i := 0; i < length; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, tmp)
		// 上面到for循环遍历完后，从queue中移除已经遍历过的元素
		// 不这么做也可以在外层for循环里的最开头将queue赋值给一个新的变量，然后将queue置为nil
		queue = queue[length:]
	}
	return result
}
