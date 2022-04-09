package leetcode

/*
	二叉树的前序遍历
*/

// 递归
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ret
}

// 迭代 + 栈
// 中序遍历 先访问根结点，然后访问左子树，然后是右子树。
// 那么我们可以使用栈，先把root加入栈，先弹出栈，然后再把root的右子树加入栈，左子树压入栈，
// 这样下一次弹出就是左子树的root，然后加入结果。然后对左子树进行相同的操作。
// 当左子树遍历完毕之后，最后弹出右子树，进行相同的操作
func preorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root == nil {
		return ret
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ret
}
