package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// SortList 排序链表
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 寻找中间节点
	// 快慢指针，快指针每次移动2步，慢指针每次移动1步，
	// 当快指针到达链表末尾时，慢指针指向的链表节点即为链表的中点。
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	// 找到中间节点后，断开链表
	rightHead := mid.Next
	mid.Next = nil

	list1 := SortList(head)
	list2 := SortList(rightHead)

	return merge(list1, list2)
}

// 合并两个有序链表
func merge(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list3 := dummy
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				list3.Next = list2
				list2 = list2.Next
			} else {
				list3.Next = list1
				list1 = list1.Next
			}
		} else if list1 != nil {
			list3.Next = list1
			list1 = nil
		} else {
			list3.Next = list2
			list2 = nil
		}
		list3 = list3.Next
	}

	return dummy.Next
}
