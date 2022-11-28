package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// GetIntersectionNode 找到两条单链表相交的起始节点
// 注意：这里的链表相交不包含两条链表十字交叉，而是指两条链表相交之后汇集到一条链表
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 根据题目的意思，链表相交之后汇集到一条链表，则交点之后走的路程一样，
	// 那么链表A走完后再继续走链表B的路程，和链表B走完后再走链表A的路程，他们走过的路程是一样的，则他们一定会在某个点相交
	// 假设链表相交的点后面的路程是c，链表A在相交前的路程是a，链表B在相交前的路程是b，则有：a+c+b+c = b+c+a+c

	// 双链表：
	if headA == nil || headB == nil {
		return nil
	}

	listA, listB := headA, headB
	for listA != nil || listB != nil {
		if listA == listB {
			return listA
		}

		if listA == nil {
			listA = headB
		} else {
			listA = listA.Next
		}

		if listB == nil {
			listB = headA
		} else {
			listB = listB.Next
		}
	}

	return nil
	// 解法一：哈希表，将链表A的元素存入哈希表，然后遍历链表B，若在哈希表找到B的元素，
	// 则表示两个链表在此处相交
	//if headA == nil || headB == nil {
	//	return nil
	//}

	//var m = map[*ListNode]struct{}{}
	//for headA != nil {
	//	m[headA] = struct{}{}
	//	headA = headA.Next
	//}

	//for headB != nil {
	//	if _, ok := m[headB]; ok {
	//		return headB
	//	}
	//	headB = headB.Next
	//}
	//return nil
}
