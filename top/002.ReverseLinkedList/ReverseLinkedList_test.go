package top

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	var list *ListNode
	for i := 1; i <= 5; i++ {
		list = addNode(list, i)
	}

	newList := ReverseLinkedList(list)
	for {
		fmt.Println(newList)
		if newList.Next == nil {
			break
		}
		newList = newList.Next
	}
}

// addNode 向链表添加节点
func addNode(list *ListNode, e int) *ListNode {
	var newNode = &ListNode{Value: e}
	if list == nil {
		list = newNode
		return list
	}

	// 为什么这样不行？
	//for list.Next != nil {
	//	list = list.Next
	//}
	//list.Next = newNode

	current := list
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return list
}
