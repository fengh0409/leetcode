package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// MergeKLists 合并k个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	// 归并排序，分治法
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	mid := length / 2
	left := MergeKLists(lists[:mid])
	right := MergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list3 := dummy
	var current *ListNode
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				current = list1
				list1 = list1.Next
			} else {
				current = list2
				list2 = list2.Next
			}
		} else if list1 == nil && list2 != nil {
			current = list2
			list2 = list2.Next
		} else if list1 != nil && list2 == nil {
			current = list1
			list1 = list1.Next
		}
		list3.Next = current
		list3 = list3.Next
	}
	return dummy.Next
}

// 方法一：顺序合并，定义一个空链表，数组中的所有链表依次与之合并，即执行多次合并两个有序链表
//func MergeKLists(lists []*ListNode) *ListNode {
//	dummy := &ListNode{}
//	tmplist := dummy
//	var klist, current *ListNode
//	for _, list := range lists {
//		for klist != nil || list != nil {
//			if klist != nil && list != nil {
//				if klist.Val < list.Val {
//					current = klist
//					klist = klist.Next
//				} else {
//					current = list
//					list = list.Next
//				}
//			} else if klist == nil && list != nil {
//				current = list
//				list = list.Next
//			} else if klist != nil && list == nil {
//				current = klist
//				klist = klist.Next
//			}
//			tmplist.Next = current
//			tmplist = tmplist.Next
//		}
//		klist = dummy.Next
//		// 重置dummy和tmplist
//		dummy = &ListNode{}
//		tmplist = dummy
//	}
//	return klist
//}
