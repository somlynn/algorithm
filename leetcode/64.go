package leetcode

/*
	最小路径和
*/

// 从左上角到右下角，到达(i,j)位置，只能从(i-1,j) 和 (i,j-1)位置过来。
// 记dp[i][j] 为 到达(i,j)最小路径和,
// 所以dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = Min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
