package leetcode

import (
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode
	// 5->6->7
	slice1 := []int{5, 6, 7}
	for _, v := range slice1 {
		l1 = addNode(l1, v)
	}

	var l2 *ListNode
	// 6->7->8->9
	slice2 := []int{6, 7, 8, 9}
	for _, v := range slice2 {
		l2 = addNode(l2, v)
	}

	// 765+9876=10641
	// 1->4->6->0->1
	l3 := AddTwoNumbers(l1, l2)
	var l3Str string
	for l3 != nil {
		l3Str += strconv.Itoa(l3.Val)
		l3 = l3.Next
	}

	if l3Str != "14601" {
		t.Fatalf("l3Str is expected 14601, but is %s", l3Str)
	}
}

// addNode 向链表添加节点
func addNode(list *ListNode, e int) *ListNode {
	var newNode = &ListNode{Val: e}
	if list == nil {
		list = newNode
		return list
	}

	current := list
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return list
}
