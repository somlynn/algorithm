package leetcode

/*
	二叉树最大宽度
*/

type AnnotateNode struct {
	node       *TreeNode
	depth, pos int
}

// 宽度优先遍历
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*AnnotateNode{{node: root, depth: 1, pos: 1}}
	curDepth, left, ans := 1, 1, 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item.node != nil {
			queue = append(queue, &AnnotateNode{node: item.node.Left, depth: item.depth + 1, pos: item.pos * 2})
			queue = append(queue, &AnnotateNode{node: item.node.Right, depth: item.depth + 1, pos: item.pos*2 + 1})
			// 此时item是这一层的最左边的节点，我们记录当前的深度和最左节点的下标
			if curDepth != item.depth {
				curDepth = item.depth
				left = item.pos
			}
			ans = Max(ans, item.pos-left+1)
		}
	}
	return ans
}
