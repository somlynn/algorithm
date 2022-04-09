package leetcode

/*
	二叉树的中序遍历
*/

// 递归
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ret = append(ret, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ret
}

// 迭代+栈
// 中序遍历，先访问左子树，根节点，右子树
// 我们需要找到最左边的节点，加入结果，然后
// 1、如果右子树，然后继续找右子树的最左边的节点
// 2、如果没有右子树，直接加入结果
func inorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, cur.Val)
		cur = cur.Right
	}
	return ret
}
