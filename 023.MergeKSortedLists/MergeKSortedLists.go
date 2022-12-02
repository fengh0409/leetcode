package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// MergeKLists 合并k个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	// TODO: 归并排序法

	dummy := &ListNode{}
	tmplist := dummy
	var klist, current *ListNode
	for _, list := range lists {
		for klist != nil || list != nil {
			if klist != nil && list != nil {
				if klist.Val < list.Val {
					current = klist
					klist = klist.Next
				} else {
					current = list
					list = list.Next
				}
			} else if klist == nil && list != nil {
				current = list
				list = list.Next
			} else if klist != nil && list == nil {
				current = klist
				klist = klist.Next
			}
			tmplist.Next = current
			tmplist = tmplist.Next
		}
		klist = dummy.Next
		// 重置dummy和tmplist
		dummy = &ListNode{}
		tmplist = dummy
	}
	return klist
}
