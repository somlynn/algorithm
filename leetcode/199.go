package leetcode

/*
	二叉树的右视图
*/
// 深度优先
func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0, 0)
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(ret) {
			ret = append(ret, root.Val)
		}
		depth++
		dfs(root.Right, depth)
		dfs(root.Left, depth)
	}
	dfs(root, 0)
	return ret
}

// 广度优先
func rightSideView2(root *TreeNode) []int {
	ret := make([]int, 0, 0)
	queue := make([]*TreeNode, 0, 0)
	if root == nil {
		return ret
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if size == 1 {
				ret = append(ret, node.Val)
			}
			size--
		}
	}
	return ret
}
