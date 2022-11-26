package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// HasCycle 环形链表，判断链表中是否有环
func HasCycle(head *ListNode) bool {
	// 思路：快慢指针，当一个链表中有环时，快慢指针都会陷入环中进行无限次移动，然后变成追及问题。
	// 每次移动，快指针移动两次，慢指针移动一次，这样移动使得快指针到慢指针到距离加一，同时慢指针到快指针到距离减一，
	// 因此，只要一直移动下去，快指针总会追上慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
