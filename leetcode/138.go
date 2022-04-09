package leetcode

/*
	复制带随机指针的链表
*/

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// 先复制，然后连接，最后拆分
func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return head
	}
	// 复制
	cur := head
	for cur != nil {
		val := cur.Val
		clone := &RandomNode{Val: val}
		clone.Next = cur.Next
		cur.Next = clone
		cur = clone.Next
	}
	// 复制random
	cur = head
	for cur != nil {
		clone := cur.Next
		if cur.Random != nil {
			clone.Random = cur.Random.Next
		}
		cur = clone.Next
	}
	// 拆分
	cur = head
	newHead := cur.Next
	for cur.Next != nil {
		clone := cur.Next
		cur.Next = clone.Next
		cur = clone
	}
	return newHead
}
