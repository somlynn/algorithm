package leetcode

/*
	移除链表元素
*/

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre != nil && pre.Next != nil {
		cur := pre.Next
		for cur != nil && cur.Val == val {
			cur = cur.Next
		}
		pre.Next = cur
		pre = pre.Next
	}
	return dummy.Next
}
