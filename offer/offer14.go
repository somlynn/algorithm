package offer

// 剪绳子

// 把一根绳子剪成多段，并且使得每段的长度乘积最大。

// 贪心

func cutRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], Max(j*(i-j), dp[j]*(i-j)))
		}
	}
	return dp[n]
}
