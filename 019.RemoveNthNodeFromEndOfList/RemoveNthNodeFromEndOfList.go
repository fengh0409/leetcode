package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type ListNode = structures.ListNode

// RemoveNthFromEnd 删除链表的倒数第 N 个结点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 解法一：快慢指针
	// 给慢指针添加个dummy节点
	dummy := &ListNode{Next: head}
	fast, slow := head, dummy
	// 快指针先遍历n个节点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 然后同时遍历快慢指针，当快指针遍历到最后一个节点时，
	// 慢指针当前所处的位置就是要删除节点的前一个节点（因为加了dummy节点）
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 将待删除节点的后驱节点赋给慢指针当前所处的位置即可
	slow.Next = slow.Next.Next
	return dummy.Next

	// 解法二：先判断长度，再删除指定节点
	//var length int
	//var current = head
	//for current != nil {
	//	current = current.Next
	//	length++
	//}

	//dummy := &ListNode{0, head}
	//current = dummy
	//for i := 0; i < length-n; i++ {
	//	current = current.Next
	//}
	//current.Next = current.Next.Next
	//return dummy.Next
}
