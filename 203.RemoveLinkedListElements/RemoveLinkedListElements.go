package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// RemoveElements 移除链表元素
func RemoveElements(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{Next: head}
	var current = dummy
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}
