package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// DeleteDuplicates 删除排序链表中的重复元素，只留下不同的数字
func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			// 把重复的值先存起来
			value := current.Next.Val
			// 发现重复的值后，遍历后续的值，如果重复，就删除重复的，遇到不重复的就终止
			for current.Next != nil {
				if current.Next.Val == value {
					current.Next = current.Next.Next
				} else {
					break
				}
			}
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}
