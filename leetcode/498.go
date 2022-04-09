package leetcode

/*
	对角线遍历
*/
// 思路：首先逐条对角线遍历，然后奇数位再翻转
func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 {
		return nil
	}
	res := make([]int, 0)
	arrList := make([][]int, 0)
	for j := 0; j < len(mat[0]); j++ {
		arr := make([]int, 0)
		row, col := 0, j
		for row < len(mat) && col >= 0 {
			arr = append(arr, mat[row][col])
			col--
			row++
		}
		arrList = append(arrList, arr)
	}

	for i := 1; i < len(mat); i++ {
		arr := make([]int, 0)
		row, col := i, len(mat[i])-1
		for row < len(mat) && col >= 0 {
			arr = append(arr, mat[row][col])
			col--
			row++
		}
		arrList = append(arrList, arr)
	}

	for i := range arrList {
		arr := arrList[i]
		if i%2 == 1 {
			for _, v := range arr {
				res = append(res, v)
			}
		} else {
			for j := len(arr) - 1; j >= 0; j-- {
				res = append(res, arr[j])
			}
		}
	}
	return res
}
