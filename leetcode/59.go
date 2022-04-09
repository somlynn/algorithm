package leetcode

/*
	螺旋矩阵 II
*/

// 和螺旋矩阵矩阵同理，不过这次是赋值
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	curNum := 1
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		for i := c1; i <= c2; i++ {
			matrix[r1][i] = curNum
			curNum++
		}
		for i := r1 + 1; i <= r2; i++ {
			matrix[i][c2] = curNum
			curNum++
		}
		if r1 != r2 {
			for i := c2 - 1; i >= c1; i-- {
				matrix[r2][i] = curNum
				curNum++
			}
		}
		if c1 != c2 {
			for i := r2 - 1; i >= r1+1; i-- {
				matrix[i][c1] = curNum
				curNum++
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return matrix
}
