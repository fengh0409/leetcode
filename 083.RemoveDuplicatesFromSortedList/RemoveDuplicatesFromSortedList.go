package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// DeleteDuplicates 删除排序链表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val != current.Next.Val {
			current = current.Next
		} else {
			current.Next = current.Next.Next
		}
	}
	return head
}
