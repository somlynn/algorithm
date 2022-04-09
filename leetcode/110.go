package leetcode

/*
	平衡二叉树
*/

// 平衡二叉树： 左子树和右子树高度差为1，并且左右子树也都是平衡树
// 递归
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return Max(depth(root.Left), depth(root.Right)) + 1
	}
	return Abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 优化,在求depth的时候就可以判断，减少递归次数
// 遍历每个节点以及子树是否是二叉树
func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(*TreeNode) (int, bool)
	depth = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		left, isBTree := depth(root.Left)
		if !isBTree {
			return 0, false
		}
		right, isBTree := depth(root.Right)
		if !isBTree {
			return 0, false
		}
		return Max(left, right) + 1, Abs(left-right) <= 1
	}
	_, isBTree := depth(root)
	return isBTree
}
