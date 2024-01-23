package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// Partition 分割链表
func Partition(head *ListNode, x int) *ListNode {
	// 定义两个新的链表，分别存储小于x的节点和大于等于x的节点
	small := &ListNode{}
	large := &ListNode{}
	smallDummy := small
	largeDummy := large
	// 遍历链表，比x小的放在small，大于等于x的放在large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	// large链表最后一个节点的Next置为nil，目的是切断原链表head的引用
	large.Next = nil
	// small链表的最后一个节点Next指向large的头结点
	small.Next = largeDummy.Next
	return smallDummy.Next
}
