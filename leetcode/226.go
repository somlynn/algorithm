package leetcode

/*
	翻转二叉树
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Left = left
	root.Right = right
	return root
}
