package leetcode

/*
	二叉树的后序遍历
*/

// 递归
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		ret = append(ret, root.Val)
	}
	dfs(root)
	return ret
}

// 迭代+ 栈
// 后序遍历 先访问左子树、右子树、根节点
// 我们拿后序遍历与先序遍历比较，
// 如果我们把后序遍历倒过来：根节点、右子树、左子树，其实是先序遍历的变形，然后reverse就行了
func postorderTraversal2(root *TreeNode) []int {
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
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverse3(ret)
	return ret
}

func reverse3(nums []int) {
	l, h := 0, len(nums)-1
	for l < h {
		nums[l], nums[h] = nums[h], nums[l]
		l++
		h--
	}
}
