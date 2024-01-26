package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// ReorderList 重排链表
func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 根据快慢指针，找到中间节点
	middle := middleNode(head)
	list2 := middle.Next
	// 从中间节点断开左右两侧链表
	middle.Next = nil
	list1 := head
	// 反转右侧链表
	list2 = reverse(list2)
	// 再合并两个链表
	merge(list1, list2)
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	var current = head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func merge(list1, list2 *ListNode) {
	for list1 != nil && list2 != nil {
		list1Next := list1.Next
		list1.Next = list2
		list1 = list1Next

		list2Next := list2.Next
		list2.Next = list1
		list2 = list2Next
	}
}
