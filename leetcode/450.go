package leetcode

/*
	删除二叉搜索树的节点
*/

// 思路：根据二叉搜索树的特性，设被删除节点node的前继节点为preNode，后继节点为nextNode
// 如果 node 没有左右孩子即为叶子节点，则直接删除 node = nil
// 如果 node 有右孩子，则用nextNode代替node,并且删除 nextNode
// 如果 node 有左孩子，则用preNode代替node，并且删除 preNode
// 关键是要找到 node 的preNode和nextNode
func getPreNode(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func getNextNode2(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val {
		root.Right = deleteNode2(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode2(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = getNextNode2(root)
			root.Right = deleteNode2(root.Right, root.Val)
		} else {
			root.Val = getPreNode(root)
			root.Left = deleteNode2(root.Left, root.Val)
		}
	}
	return root
}
