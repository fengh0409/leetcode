package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// ReverseKGroup K 个一组翻转链表
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		tail := prev
		// 每次遍历k个节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				break
			}
		}
		if tail == nil {
			break
		}
		// 定义待翻转的链表的第一个节点
		start := prev.Next
		// 定义下一部分的第一个节点
		next := tail.Next
		// 切断待翻转链表的最后一个节点
		tail.Next = nil
		// 翻转该部分链表，翻转后的链表就变成了第一个节点
		prev.Next = reverse(start)
		// 下一个待翻转的链表的第一个节点就是之前保存的tail的下一个节点
		start.Next = next

		// start变成头节点
		prev = start
		// 头尾节点一样
		tail = prev
	}
	return dummy.Next
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
