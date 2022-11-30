package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// IsPalindrome 判断链表是否为回文链表
func IsPalindrome(head *ListNode) bool {
	// 思路：找到链表的中间位置，将右半部分链表反转，然后同时遍历左右链表，比较每个元素值，
	// 元素全部相同则为回文链表，否则不是回文链表
	if head == nil {
		return true
	}

	left := findMiddleNode(head)
	rightReverseList := reverseLinkedList(left.Next)
	p1, p2 := head, rightReverseList
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

// 通过快慢指针找到中间节点
func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
