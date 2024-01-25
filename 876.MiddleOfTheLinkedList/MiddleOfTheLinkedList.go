package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// MiddleNode 链表的中间节点
func MiddleNode(head *ListNode) *ListNode {
	// 链表节点数为0时，这种解法不通过
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast.Next != nil && fast.Next.Next == nil {
		return slow.Next
	}
	return slow

	// 或者
	//fast, slow := head, head
	//for fast != nil && fast.Next != nil {
	//	fast = fast.Next.Next
	//	slow = slow.Next
	//}
	//return slow
}
