package leetcode

/*
	二叉树的锯齿形层序遍历
*/

// 时间复杂度: O(N) 空间复杂度: O(N)
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	flag := true
	for len(queue) > 0 {
		size := len(queue)
		list := make([]int, size)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if flag {
				list[len(list)-size] = node.Val
			} else {
				list[size-1] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		flag = !flag
		ret = append(ret, list)
	}
	return ret
}
