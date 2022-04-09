package offer

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BuildListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		node := &ListNode{num, nil}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

func BuildTreeNode(nums []string) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var dfs func(int) *TreeNode
	dfs = func(i int) *TreeNode {
		if i >= len(nums) || nums[i] == "null" {
			return nil
		}
		val, _ := strconv.Atoi(nums[i])
		root := &TreeNode{Val: val}
		root.Left = dfs(i*2 + 1)
		root.Right = dfs(i*2 + 2)
		return root
	}
	root := dfs(0)
	return root
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	newHead := reverseList(next)
	next.Next = head
	return newHead
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		next := l1.Next
		l1.Next = mergeTwoLists(next, l2)
		return l1
	} else {
		next := l2.Next
		l2.Next = mergeTwoLists(l1, next)
		return l2
	}
}

func reverse3(nums []int) {
	l, h := 0, len(nums)-1
	for l < h {
		nums[l], nums[h] = nums[h], nums[l]
		l++
		h--
	}
}
