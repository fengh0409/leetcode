package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// SwapPairs 两两交换链表中的节点
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	list := dummy
	for list.Next != nil && list.Next.Next != nil {
		node1 := list.Next
		node2 := list.Next.Next
		// 没什么花活，就是两两交换，原本第二个节点的子节点成为第一个节点的子节点
		node1.Next = node2.Next
		// 原本第一个节点成为第二个节点的子节点
		node2.Next = node1
		// 原本第二个节点成为头结点
		list.Next = node2
		// 下次遍历开始的头结点
		list = node1
	}
	return dummy.Next
}
