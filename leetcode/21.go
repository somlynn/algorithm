package leetcode

/*
	合并两个有序链表
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		next := l1.Next
		l1.Next = mergeTwoLists(next, l2)
		return l1
	} else {
		next := l2.Next
		l2.Next = mergeTwoLists(l1, next)
		return l2
	}
}
