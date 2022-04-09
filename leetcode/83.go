package leetcode

/*
	删除排序链表中的重复元素
*/

// 递归
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	for cur != nil && cur.Val == head.Val {
		cur = cur.Next
	}
	head.Next = deleteDuplicates(cur)
	return head
}
