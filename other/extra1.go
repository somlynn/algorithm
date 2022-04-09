package other

/*
	给定一个奇数位升序，偶数位降序的链表，将其重新排序。
*/

func sortOddEvenList(head *ListNode) *ListNode {
	oldHead, newHead := splitList(head)
	newHead = reverseList(newHead)
	return mergeTwoLists(oldHead, newHead)
}

// 先拆分
func splitList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}
	newHead := head.Next
	old, cur := head, head.Next

	for cur.Next != nil {
		old.Next = cur.Next
		old = cur
		cur = cur.Next
	}
	old.Next = nil
	return head, newHead
}
