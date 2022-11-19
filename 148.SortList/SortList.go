package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// SortList 排序链表
func SortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// ??
	if head.Next == tail {
		head.Next = nil
		return head
	}

	// 快慢指针，快指针每次移动2步，慢指针每次移动1步，
	// 当快指针到达链表末尾时，慢指针指向的链表节点即为链表的中点。
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		// 快指针每次移动2步
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}

	return dummy.Next
}
