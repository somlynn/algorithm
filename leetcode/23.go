package leetcode

/*
	合并K个升序链表
*/

// 分治+合并两个有序链表
// 时间复杂度 O(n*logk) n表示所有链表长度总和，k表示链表的个数
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l, h := 0, len(lists)
	m := l + (h-l)/2
	left := mergeKLists(lists[l:m])
	right := mergeKLists(lists[m:])
	return mergeTwoLists(left, right)
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	var dfs func(lists []*ListNode, l, h int) *ListNode
	dfs = func(lists []*ListNode, l, h int) *ListNode {
		if l == h {
			return lists[l]
		}
		m := l + (h-l)/2
		left := dfs(lists, l, m)
		right := dfs(lists, m+1, h)
		return mergeTwoLists(left, right)
	}
	return dfs(lists, 0, len(lists)-1)
}
