package leetcode

/*
	买卖股票的最佳时机
*/

// 时间复杂度: O(N) 空间复杂度: O(1)
// minPrice记录最小值，maxProfit记录最大利润。
// 从左向右遍历元素，并且记录最小值，
// 当前元素与最小值比较，差值大于maxProfit 则更新maxProfit
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if p := prices[i] - minPrice; p > maxProfit {
			maxProfit = p
		}
	}
	return maxProfit
}
