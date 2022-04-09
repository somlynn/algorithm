package other

//  二叉树的下一个结点

// 给定二叉树其中的一个结点，请找出其中序遍历顺序的下一个结点并且返回。
//注意，树中的结点不仅包含左右子结点，而且包含指向父结点的指针。

type TreeLinkNode struct {
	Left   *TreeLinkNode
	Right  *TreeLinkNode
	Parent *TreeLinkNode
}

// 1、node 有右子树，则 node 下一个结点就是 右子树的最左边的结点
// 2、node 没有右子树
//		a. 如果node是左结点，node的下一个结点就是 parent
// 		b. 如果node是右结点，node的下一个结点就是 if parent == parent.parent.Left , next = parent.parent
func getNextNode(node *TreeLinkNode) *TreeLinkNode {
	if node.Right != nil {
		right := node.Right
		for right.Left != nil {
			right = right.Left
		}
		return right
	} else {
		for node.Parent != nil {
			parent := node.Parent
			if parent.Left == node {
				return parent
			}
			node = node.Parent
		}
	}
	return nil
}
