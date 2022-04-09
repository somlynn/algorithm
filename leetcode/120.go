package leetcode

/*
	三角形最小路径和
*/

// 动态规划
// 我们用 dp[i][j] 表示从三角形顶部走到位置 (i, j) 的最小路径和。这里的位置 (i, j)指的是三角形中第 i 行第 j 列（均从 0 开始编号）的位置。
// 由于每一步只能移动到下一行的相邻节点，因此想走到位置(i,j)，上一个步就只能在位置(i-1,j-1)或者位置(i-1,j)
// 我们在这个两个位置取最小的 dp[i][j] = min(dp[i-1][j-1],dp[i-1][j])+nums[i][j]
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n, n)
	for i := range dp {
		dp[i] = make([]int, n, n)
	}
	// 初始化 三角形的顶点和 两边 即 j==0 只能从(i-1,0)到(i,0)和 j==i只能从(i-1,i-1)到(i,i)
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = Min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}
	// 然后比较底部边 dp[n-1][j] 的值 取最小值
	return dp[0][0]
}

// 优化一，从上而下求，底边会有很多扇出，我们需要比较
// 我们可以从下而上求

func minimumTotal3(triangle [][]int) int {
	return 0
}

// 动态规划+空间优化
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = Min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
