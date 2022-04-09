package leetcode

/*
	删除链表中的节点
*/

func deleteNode(node *ListNode) {
	next := node.Next
	towNext := next.Next
	node.Val = next.Val
	node.Next = towNext
	next.Next = nil
}
