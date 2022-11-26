package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// DetectCycle 环形链表II，返回链表开始入环的第一个节点，如果链表无环，则返回 nil
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	// ???
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast

	// 解法一：哈希表
	//var m = map[*ListNode]struct{}{}
	//for head != nil {
	//	if _, ok := m[head]; ok {
	//		return head
	//	}
	//	m[head] = struct{}{}
	//	head = head.Next
	//}
	//return nil
}
