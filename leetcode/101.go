package leetcode

/*
	对称二叉树
*/

// 递归
func isSymmetric(root *TreeNode) bool {
	return isSymmetricTwo(root, root)
}

// 判断两颗数是不是对称
func isSymmetricTwo(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSymmetricTwo(p.Left, q.Right) && isSymmetricTwo(p.Right, q.Left)
}

// 迭代
func isSymmetric2(root *TreeNode) bool {
	p, q := root, root
	queue := make([]*TreeNode, 0)
	queue = append(queue, p, q)
	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil || p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Right, p.Right, q.Left)
	}
	return true
}
