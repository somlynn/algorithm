package leetcode

/*
	从前序与中序遍历序列构造二叉树
*/

// 前序遍历：preorder[i] 是root,preorder[i+1,j] 是左子树的所有节点，inorder[j+1,n]是右子树的所有的节点
// 但是我们不知道 左右子树的长度，借助中序遍历我们可以得出，左子树的长度知道了，剩下的就是右子树了
// 中序遍历：inorder[i] 是root ，则 [0,i-1]是左子树，size= i,
// 那么假如preorder[i]是root,左子树preorder[i+,i+size],右子树preorder[i+size+1,n]
// 然后在递归左右的前序序列，形成左右子树，
// 即root.Left = dfs(preorder[i,i+size]),root.Right=dfs(preorder[i+size+1,n])

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	var dfs func([]int, int, int, int) *TreeNode
	dfs = func(pre []int, preL int, preR int, inL int) *TreeNode {
		if preL > preR {
			return nil
		}
		root := &TreeNode{Val: pre[preL]}
		index := indexMap[root.Val]
		leftSize := index - inL
		root.Left = dfs(pre, preL+1, preL+leftSize, inL)
		root.Right = dfs(pre, preL+leftSize+1, preR, index+1)
		return root
	}
	return dfs(preorder, 0, len(preorder)-1, 0)
}
