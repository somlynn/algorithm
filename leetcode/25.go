package leetcode

/*
	K个一组翻转链表
*/

// 递归 时间复杂度: O(N) 空间复杂度: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	// 首先找到第k个节点，可能链表的长度小于k，遍历时要判空
	// 达到第k个节点，需要移动k-1次
	for i := 1; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	// 遍历结束后，如果cur==nil，说明链表的长度小于等于k，直接返回
	if cur == nil {
		return head
	}
	next := cur.Next
	cur.Next = nil
	newHead := reverseList(head)
	nextHead := reverseKGroup(next, k)
	head.Next = nextHead
	return newHead
}
