package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// OddEvenList 奇偶链表
func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 奇节点的头
	odd := head
	// 偶节点
	even := odd.Next
	// 偶节点的头
	evenHead := even
	// 遍历偶节点
	for even != nil && even.Next != nil {
		// 奇节点的下一个节点就是指向偶节点的下一个节点
		odd.Next = even.Next
		odd = odd.Next
		// 同理，偶节点的下一个节点就是指向奇节点的下一个节点
		even.Next = odd.Next
		even = even.Next
	}
	// 奇节点的最后一个节点指向偶节点的头
	odd.Next = evenHead
	return head
}

// 解法二：类似086题分割链表，先分别将奇偶链表放到两个链表，再合并
//func OddEvenList(head *ListNode) *ListNode {
//	var leftList = &ListNode{}
//	var rightList =&ListNode{}
//	leftDummy := leftList
//	rightDummy := rightList
//	var n = 1
//	for head != nil {
//		if n%2 > 0 {
//			leftList.Next = head
//			leftList = leftList.Next
//		} else {
//			rightList.Next = head
//			rightList = rightList.Next
//		}
//		head = head.Next
//		n++
//	}
//	rightList.Next = nil
//	leftList.Next = rightDummy.Next
//	return leftDummy.Next
//}
