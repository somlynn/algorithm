package leetcode

/*
	求根节点到叶节点数字之和
*/

// 递归
func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, num int) int {
		if root == nil {
			return 0
		}
		num = num*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return num
		}
		return dfs(root.Left, num) + dfs(root.Right, num)
	}
	return dfs(root, 0)
}
