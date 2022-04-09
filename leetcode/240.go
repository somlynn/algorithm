package leetcode

/*
	搜索二维矩阵 II
*/

// 从右上角开始搜索
// 时间复杂度: O(r+c) 空间复杂度O(1)
func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	i, j := 0, c-1
	for i < r && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
