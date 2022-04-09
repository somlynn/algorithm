package other

/*
	圆环上有10个点，编号为0~9。从0点出发，每次可以逆时针和顺时针走一步，问走n步回到0点共有多少种走法。
*/

// 如果你之前做过leetcode的70题爬楼梯，则应该比较容易理解：
// 走n步到0的方案数=走n-1步到1的方案数+走n-1步到9的方案数。
// dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
// 公式之所以取余是因为j-1或j+1可能会超过圆环0~9的范围
// dp[i][j]定义为走 i 到 j 这个位置的方案数。
func backToOrigin(n int) int {
	length := 10
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, length)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < length; j++ {
			dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
		}
	}
	return dp[n][0]
}
