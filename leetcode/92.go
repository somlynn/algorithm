package leetcode

/*
	反转链表 II
*/

// 递归
// 可以先考虑反转[1,n]的情况。然后递归
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		cur := head
		for i := 1; i < right; i++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		newHead := reverseList(head)
		head.Next = next
		return newHead
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 一次遍历+头插法
// 找第几个节点的技巧：找第n个节点，需要移动n-1次，即循环n-1次
// 注意头插法时，需要循环right-left次。
// 很多时候我们需要知道前一个节点，但是head没有前一个节点，这里有个小技巧：设置虚节点可以方便我们解答问题

// 头插法
// 反转[m,n] 中间需要反转的长度为 n-m+1
// 利用头插法，举例：pre->1->2->3->4->post
// 1、pre->2->1->3->4->post
// 2、pre->3->2->1->4->post
// 3、pre->4->3->2->1->post
// 设size=n-m+1,需要size-1次头插法，因此需要循环size-1次,即循环n-m次
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 1; i <= right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
