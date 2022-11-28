package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// GetIntersectionNode 找到两条单链表相交的起始节点
// 注意：这里的链表相交不包含两条链表十字交叉，而是指两条链表相交之后汇集到一条链表
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双链表：

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
