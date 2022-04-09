package leetcode

/*
	重排链表
*/

// 把链表分成两半，反转后面的链表，然后交替重新连接
// 链表找中间节点的技巧：快慢指针，
// 采用slow = head,fast = head.Next的方式，对于偶数能找到前一半最后的一个节点，对于奇数正好找到中间节点
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	s := slow.Next
	slow.Next = nil
	s = reverseList(s)
	cur := head
	other := s
	for other != nil {
		next := cur.Next
		cur.Next = other
		other = next
		cur = cur.Next
	}
}
