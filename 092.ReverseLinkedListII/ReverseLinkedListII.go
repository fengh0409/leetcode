package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// ReverseBetween 反转链表II
// 1->2->3->4->5
// 1->4->3->2->5
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	leftBefore := dummy
	// left的前一个节点
	for i := 0; i < left-1; i++ {
		leftBefore = leftBefore.Next
	}
	// left所在的节点
	leftNode := leftBefore.Next

	// right所在的节点
	rightNode := leftBefore
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	// right的下一个节点
	rightNext := rightNode.Next

	// 从left和right节点处切断链表
	leftBefore.Next = nil
	rightNode.Next = nil

	// 反转链表
	reverseLinkedList(leftNode)

	// 反转完成后，原来的右节点在左侧，将其连接到left的前一个节点
	leftBefore.Next = rightNode
	// 原来的左节点在右侧，将right的下一个节点连接到其后面
	leftNode.Next = rightNext

	return dummy.Next
}

func reverseLinkedList(head *ListNode) {
	var prev *ListNode
	var current = head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
}
