package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// MergeTwoSortedLists 合并两个有序链表
func MergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 解法一：
	// 类似合并两个有序数组，同时遍历两条链表，比较数值大小
	var dummy = &ListNode{}
	var list3 = dummy

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

	// 解法二：递归
	//if list1 == nil {
	//	return list2
	//}
	//if list2 == nil {
	//	return list1
	//}
	//if list1.Val < list2.Val {
	//	list1.Next = MergeTwoSortedLists(list1.Next, list2)
	//	return list1
	//}
	//list2.Next = MergeTwoSortedLists(list1, list2.Next)
	//return list2
}
