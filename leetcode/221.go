package leetcode

/*
	最大正方形
*/

// 动态规划
// 总体思想：设dp[i][j] 表示 以matrix[i-1][j-1]为右下角正方形的最大边长
//
func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// matrix[i][j] == '0' 以matrix[i-1][j-1]为右下角形成不了正方形 dp[i][j] = 0
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					// dp[i][j]的值由其上方、左方和左上方的三个相邻位置的dp值决定,木桶效应
					// 当前位置的元素值等于三个相邻位置的元素中的最小值加 1
					dp[i][j] = Min(Min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
				}
				maxSide = Max(maxSide, dp[i][j])
			}
		}
	}
	return maxSide * maxSide
}

// 为了方便统计我们 从(1,1)开始
func maximalSquare2(matrix [][]byte) int {
	maxSide := 0
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			//
			if matrix[i-1][j-1] == '1' {
				// dp[i][j]的值由其上方、左方和左上方的三个相邻位置的dp值决定,木桶效应
				// 当前位置的元素值等于三个相邻位置的元素中的最小值加 1
				dp[i][j] = Min(Min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
				maxSide = Max(maxSide, dp[i][j])
			}
		}
	}
	return maxSide * maxSide
}
