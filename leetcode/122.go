package leetcode

/*
	买卖股票的最佳时机II
*/

// 时间复杂度: O(N) 空间复杂度: O(1)
// 允许多次买卖，最大利润就是 所有大于0差值的和,即只增不减便是最大利润。
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - prices[i-1]; diff > 0 {
			maxProfit += diff
		}
	}
	return maxProfit
}
