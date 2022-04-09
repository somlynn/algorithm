package leetcode

/*
	二叉树的直径
*/

// 注意这个题目求得是直径是 边数 = 节点数-1
// 经过 root 的直径为 左子树的深度+右子树的深度+本身
// 然后递归求经过所有节点的直径，然后比较就最大的。
func diameterOfBinaryTree(root *TreeNode) int {
	maxLen := 0
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left)
		right := depth(root.Right)
		maxLen = Max(maxLen, left+right+1)
		return Max(left, right) + 1
	}
	depth(root)
	return maxLen - 1
}
