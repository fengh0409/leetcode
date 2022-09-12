package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 两数相加，两个链表相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, current *ListNode
	var carry int

	// 时间复杂度 O(max(m,n))，max(m,n)为两个链表最长的值
	// 同时遍历两个链表，都为空时则退出
	for l1 != nil || l2 != nil {
		// l1和l2的链表长度可能不一致，当只剩一个链表在遍历时，另一个链表的值为0
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		// 将两个链表同一位置的值相加，大于10则往后进1位，由于两数相加的和大于10时，下一次求和需要加上进位carry，所以求和实际是三数之和，没有进位时carry=0
		sum := l1Val + l2Val + carry
		// 三数相加取余，即为该节点的值
		val := sum % 10
		// 三数相加取整，即为进位，进位存起来，下一次遍历会用到
		carry = sum / 10

		// 第一次遍历
		if l3 == nil {
			l3 = &ListNode{Val: val}
			// 用l3给current赋值，通过current改变l3的值
			current = l3
		} else {
			current.Next = &ListNode{Val: val}
			// 最后节点向后移位
			current = current.Next
		}
	}

	// 若遍历结束后还剩一个进位，则最后再加一个节点
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return l3
}

// 以下为错误解法，这种方式会发生求和溢出的情况
//func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	// 先反转链表
//	l1 = reverseLinkedList(l1)
//	l2 = reverseLinkedList(l2)
//
//	// 然后将链表的值拼接成字符串，再转换为整数
//	var l1ValueStr, l2ValueStr string
//	for l1 != nil {
//		l1ValueStr += strconv.Itoa(l1.Val)
//		l1 = l1.Next
//	}
//
//	for l2 != nil {
//		l2ValueStr += strconv.Itoa(l2.Val)
//		l2 = l2.Next
//	}
//
//	l1ValueInt, _ := strconv.Atoi(l1ValueStr)
//	l2ValueInt, _ := strconv.Atoi(l2ValueStr)
//
//	// 再求和
//	l3ValueInt := l1ValueInt + l2ValueInt
//	l3ValueStr := strconv.Itoa(l3ValueInt)
//
//	// 再将每个数字倒序存入链表
//	var l3 *ListNode
//	parts := strings.Split(l3ValueStr, "")
//	for i := len(parts) - 1; i >= 0; i-- {
//		v, _ := strconv.Atoi(parts[i])
//		l3 = addNode(l3, v)
//	}
//
//	return l3
//}
//
//func reverseLinkedList(list *ListNode) *ListNode {
//	if list == nil {
//		return nil
//	}
//
//	var prev *ListNode
//	current := list
//	for current != nil {
//		next := current.Next
//		current.Next = prev
//		prev = current
//		current = next
//	}
//
//	return prev
//
//}
//
