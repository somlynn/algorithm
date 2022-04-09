package leetcode

/*
	零钱兑换II   可以凑成总金额的硬币组合数
*/

// 转化为背包问题
// 记dp[i] 表示当前背包容量为i时，可以凑成i的硬币组合数
// 初始化：dp[0]=1，当背包容量为 0 时，只有一种情况满足条件，就是一个硬币也不选
// 假设有硬币: x1,x2,x3,
// dp[i] = dp[i-x1] + dp[i-x2] + dp[i-x3]
// dp[j] = dp[j-x1] + dp[i-x2] + dp[i-x3]
// 转化为变量遍历
// for _ x := coins
// 		for i := x ;i <= amount ; i++
//         dp[i] += dp[i-x]  // 因为 i-x >=0 ，所以i:=x 作为起点
//
func change(amount int, coins []int) int {
	dp := make([]int, amount+1, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i < amount+1; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
