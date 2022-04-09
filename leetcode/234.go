package leetcode

/*
	回文链表
*/

//按之前技巧找到中间节点，分成2半，
// 反转右一半, len(p1) >= len(p2),所以我们以判断p2 !=nil 作为遍历条件
// 如果中间有不等，则返回false，否则返回true
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := getMiddleNode(head)
	right := reverseList(mid.Next)
	p1, p2 := head, right
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func getMiddleNode(head *ListNode) *ListNode {
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
