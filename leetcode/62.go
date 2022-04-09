package leetcode

/*
	不同路径
*/

// 动态规划
// 因为机器人只能向右和向下移动 dp[i][j]为移动第i，j个格子的不同路径个数
// 那么dp[i][j] = dp[i-1][j]+dp[i][j-1]
// 首先要初始化，因为机器人到达dp[i][0] 和 dp[0][j]都是1种路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
