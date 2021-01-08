package top

// ListNode 链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

// ReverseLinkedList 反转链表
// 1->2->3->4->5->nil
// nil<-1<-2<-3<-4<-5
func ReverseLinkedList(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}

	// 迭代法：时间复杂度为O(n)
	// 声明新的链表
	var prev *ListNode
	current := list
	for current != nil {
		// 把下一个节点存起来
		next := current.Next
		// 反转当前节点的指向，将当前节点的下一个节点指向前一个节点（也是新的链表prev），比如第一次遍历，将第一个节点指向 nil
		current.Next = prev
		// 因为当前节点的下一个节点指向前一个节点，所以当前节点 current 变为了 prev
		// nil<-1  2->3->4->5，这里的1就是prev，2就是current
		// nil<-1<-2  3->4->5，这里的2就是prev，3就是current
		prev = current
		// 前面赋值完成后，当前节点需要向后移动一位，即 current = current.Next，但是不能直接这样赋值，因为 current.Next 在前面已经赋给 prev 了，所以需要一个临时变量 next 把当前节点的下一个节点提前存起来
		current = next
	}

	return prev

	// 递归
	//if list == nil || list.Next == nil {
	//	return list
	//}
	//newList := ReverseLinkedList(list.Next)
	// 当list.Next 为最后一个节点时，将其指向的节点反转，即改为指向前一个节点，再将前一个节点原有的指向清空
	// 例如：当list.Next = 5时，很明显此时list = 4，原来是4指向5，5指向nil，变为5指向4，4指向nil，
	// 即 1->2->3->4->5->nil
	// 变 1->2->3->4<-5
	//list.Next.Next = list
	//list.Next = nil
	//return newList
}
