package leetcode

/*
	删除链表的倒数第 N 个结点
*/
// 快慢指针
// 增加虚节点，找到倒数第N个节点的前一个节点。这样如果删除第一节点也好处理
// 如果要找到倒数第N个节点，需要fast比slow多移动N次
// 但是删除节点，我们需要找到倒数第N+1个节点，需要fast比slow多移动N+1次
// 时间复杂度:O(N) , 空间复杂度: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	target := slow.Next
	slow.Next = target.Next
	target.Next = nil
	return dummy.Next
}
