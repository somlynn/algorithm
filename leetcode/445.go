package leetcode

/*
	两数相加 II
*/
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0, 0), make([]int, 0, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	i, j := len(stack1), len(stack2)
	var carry int
	dammy := &ListNode{}
	for i >= 0 || j >= 0 {
		a, b := 0, 0
		if i >= 0 {
			a = stack1[i]
		}
		if j >= 0 {
			b = stack2[j]
		}
		sum := a + b + carry
		val := sum % 10
		carry = sum / 10
		node := &ListNode{Val: val}
		next := dammy.Next
		node.Next = next
		dammy.Next = node
		i--
		j--
	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		next := dammy.Next
		node.Next = next
		dammy.Next = node
	}
	return dammy.Next
}
