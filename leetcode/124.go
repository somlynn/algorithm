package leetcode

import "math"

/*
	二叉树中的最大路径和
*/

// 递归
// 总体思路：遍历二叉树的每个节点，经过 node(i)最大路径和记为sum(i)
//		1.sum(node(i)) = sum(node(i).Left) + sum(node(i).Right)+node(i).val.如果left和right小于0,则不作为选取节点
// 		2.maxSum=Max(sum([0~n])
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := Max(depth(root.Left), 0)
		right := Max(depth(root.Right), 0)
		sum := left + right + root.Val
		maxSum = Max(maxSum, sum)
		return Max(left, right) + root.Val
	}
	depth(root)
	return maxSum
}
