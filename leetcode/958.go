package leetcode

/*
	二叉树的完全性检验
*/

// 对于完全二叉树 如果根节点编号为code 那么左子节点编号2*code，右子节点编号为2*code+1
// 判断是不是完全二叉树，即节点的个数 size 是不是等于 maxCode
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	size, maxCode := 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, code int) {
		if root == nil {
			return
		}
		size++
		maxCode = Max(maxCode, code)
		dfs(root.Left, code*2)
		dfs(root.Right, code*2+1)
	}
	dfs(root, 1)
	return size == maxCode
}

func isCompleteTree2(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	root.Val = 1
	queue = append(queue, root)
	var count int
	var maxCode int
	for len(queue) > 0 {
		size := len(queue)
		count += size
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			maxCode = node.Val
			if node.Left != nil {
				node.Left.Val = 2 * node.Val
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 1
				queue = append(queue, node.Right)
			}
			size--
		}
	}
	return maxCode == count
}
