package leetcode

/*
	删除排序链表中的重复元素 II
*/

// 当head为nil或者只有一个节点，直接return
// 当大于等于两个节点时，判断next节点是否和head的值相等，
// 如果相等，for循环知道找到下一个不相等的节点，然后递归求解 next
// 如果不想等, head.Next指向下一个节点递归的结果。
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	if next.Val == head.Val {
		for next != nil && next.Val == head.Val {
			next = next.Next
		}
		return deleteDuplicates2(next)
	} else {
		head.Next = deleteDuplicates2(next)
		return head
	}
}
