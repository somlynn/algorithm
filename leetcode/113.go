package leetcode

/*
	路径总和II
*/

// 深度优先搜索
// 从root开始深度优先搜索，
// 当是叶子节点并且路径和为targetSum 即 node.Left == nil && node.Right == nil && sum == 0 加入结果集
// 注意因为是共用一个path切片，可能共用一个底层数组，所以要copy在加入到结果集，不然会导致后面的改变修改原数组结果。
// 左子树递归结束后，要回退一个元素，然后进行右子树的递归，所以 要path = path[:len(path)-1]
func pathSum(root *TreeNode, targetSum int) [][]int {
	path := make([]int, 0, 0)
	res := make([][]int, 0, 0)
	var backtrack func(*TreeNode, int)
	backtrack = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			res = append(res, copyPath)
			return
		}
		backtrack(node.Left, sum)
		backtrack(node.Right, sum)
		path = path[:len(path)-1]
	}
	backtrack(root, targetSum)
	return res
}
