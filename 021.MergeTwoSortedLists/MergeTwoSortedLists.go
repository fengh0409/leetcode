package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists 合并两个有序链表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 解法一：
	// 类似合并两个有序数组，同时遍历两条链表，比较数值大小
	var list3 = &ListNode{}
	var dummy = list3

	for list1 != nil || list2 != nil {
		var current *ListNode
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				current = list2
				list2 = list2.Next
			} else {
				current = list1
				list1 = list1.Next
			}
		} else if list1 != nil && list2 == nil {
			current = list1
			list1 = list1.Next
		} else if list1 == nil && list2 != nil {
			current = list2
			list2 = list2.Next
		}

		list3.Next = current
		list3 = list3.Next
	}

	return dummy.Next
}
