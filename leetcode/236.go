package leetcode

/*
	二叉树的最近公共祖先
*/

//	以root为跟节点的树，p,q 两个节点存在三种情况
//  1、都在左子树上
//  2、都在右子树上
//  3、分别在左右子树上
// 递归终点就是 nil 节点，p,q节点

// 分别递归左右子树，如果left==nil，说明结果在递归的右子树上，
// 如果right == nil,说明结果在递归的左子树上
// 如果都不为nil，left ,right 即为p,q这连个节点，那么root就是公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
