package leetcode

import "math"

/*
	验证二叉树搜索树
*/
// 二叉搜索树:
// 如果左子树不为空，左子树的所有节点的值都小于根节点的值，如果右子树不为空，所有的右节点的值都大于根节点的值
// 递归
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(*TreeNode, int, int) bool
	helper = func(root *TreeNode, low int, high int) bool {
		if root == nil {
			return true
		}
		if root.Val <= low || root.Val >= high {
			return false
		}
		return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 迭代，中序遍历
// 当前节点一定比中序遍历的前一个节点大，即中序遍历是递增的。
// 记pre为当前的节点的前一个节点。
func isValidBST2(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return true
	}
	cur := root
	pre := math.MinInt64
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val <= pre {
			return false
		}
		pre = cur.Val
		cur = cur.Right
	}
	return true
}
