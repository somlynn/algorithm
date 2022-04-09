package leetcode

/*
	螺旋矩阵
*/

// 遍历外层之后，又是新的矩阵，所以遍历外层的逻辑是一样的。
// 注意两个特殊情况,只需要遍历一次就行，从左往右，从上到下。
// 1、当只有一行时，即r1==r2
// 2、当只有一列时，即c1==c2
func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	if len(matrix) == 0 {
		return ret
	}
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		for i := c1; i <= c2; i++ {
			ret = append(ret, matrix[r1][i])
		}
		for i := r1 + 1; i <= r2; i++ {
			ret = append(ret, matrix[i][c2])
		}
		if r1 != r2 {
			for i := c2 - 1; i >= c1; i-- {
				ret = append(ret, matrix[r2][i])
			}
		}
		if c1 != c2 {
			for i := r2 - 1; i >= r1+1; i-- {
				ret = append(ret, matrix[i][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return ret
}
