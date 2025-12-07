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
func TestAddTwoNumbers_EmptyLists(t *testing.T) {
	var l1, l2 *ListNode

	l3 := AddTwoNumbers(l1, l2)
	if l3 != nil {
		t.Fatalf("Expected nil, but got %v", l3)
	}
}

func TestAddTwoNumbers_OneEmptyList(t *testing.T) {
	var l1 *ListNode
	l2 := &ListNode{Val: 1}

	l3 := AddTwoNumbers(l1, l2)
	if l3.Val != 1 || l3.Next != nil {
		t.Fatalf("Expected 1, but got %v", l3)
	}
}

func TestAddTwoNumbers_SameLengthLists(t *testing.T) {
	var l1 *ListNode
	slice1 := []int{2, 4, 3}
	for _, v := range slice1 {
		l1 = addNode(l1, v)
	}

	var l2 *ListNode
	slice2 := []int{5, 6, 4}
	for _, v := range slice2 {
		l2 = addNode(l2, v)
	}

	// 342 + 465 = 807
	// 7->0->8
	l3 := AddTwoNumbers(l1, l2)
	expected := []int{7, 0, 8}
	for _, v := range expected {
		if l3 == nil || l3.Val != v {
			t.Fatalf("Expected %d, but got %v", v, l3)
		}
		l3 = l3.Next
	}
}

func TestAddTwoNumbers_DifferentLengthLists(t *testing.T) {
	var l1 *ListNode
	slice1 := []int{9, 9, 9, 9, 9, 9, 9}
	for _, v := range slice1 {
		l1 = addNode(l1, v)
	}

	var l2 *ListNode
	slice2 := []int{9, 9, 9, 9}
	for _, v := range slice2 {
		l2 = addNode(l2, v)
	}

	// 9999999 + 9999 = 10009998
	// 8->9->9->9->0->0->0->1
	l3 := AddTwoNumbers(l1, l2)
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
	for _, v := range expected {
		if l3 == nil || l3.Val != v {
			t.Fatalf("Expected %d, but got %v", v, l3)
		}
		l3 = l3.Next
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
