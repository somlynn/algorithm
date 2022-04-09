package leetcode

/*
	零钱兑换
*/

// 记dp[i] 表示 凑成i元x需要最少的硬币个数.
// 将设现在有3种硬币,x1,x2,x3，dp[i] = min(dp[i-x1],dp[i-x2],dp[i-x3])+1
//  对于没有任何一种硬币组合组成总额怎么办呢？
// 这里有一个小技巧：dp 初始化一个很大的数，如果dp[i]一直都没有更新，说明没有任何组合满足。
// 这里dp的长度设置的是amount+1 ,为了让下标i从1开始，dp[amount]即为结果
// 对于金额amount 它的组合金币数最多为 amount ,我们初始化dp为amount+1即可，
func coinChange(coins []int, amount int) int {
	size := amount + 1
	dp := make([]int, size)
	for i := 1; i < len(dp); i++ {
		dp[i] = size
	}
	dp[0] = 0
	for i := 1; i < size; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == size {
		return -1
	}
	return dp[amount]
}
