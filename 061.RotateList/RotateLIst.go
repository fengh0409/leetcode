package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// RotateRight 旋转链表
func RotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	add := n - k%n
	if add == n {
		return head
	}

	tail.Next = head
	for add > 0 {
		tail = tail.Next
		add--
	}
	newHead := tail.Next
	tail.Next = nil
	return newHead
}

// 错误解法
//func RotateRight(head *ListNode, k int) *ListNode {
//	if head == nil || head.Next == nil || k == 0 {
//		return head
//	}
//
//	dummy := &ListNode{Next: head}
//	current := dummy
//	var tail = head
//	for current.Next != nil {
//		current = current.Next
//		tail = current
//	}
//	tail.Next = head
//	current = head
//	for i := 0; i < k; i++ {
//		current = current.Next
//	}
//	dummy.Next = current.Next
//	current.Next = nil
//
//	return dummy.Next
//}
