package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var list *ListNode
	for i := 1; i <= 5; i++ {
		list = addNode(list, i)
	}
	newList := RemoveNthFromEnd(list, 2)
	fmt.Println(*newList)

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
