package leetcode

import "math"

/*
	买卖股票的最佳时机III
*/

// 如果产生了第二笔交易那第一笔一定是完成的了。整个交易一共有四种状态
// 第一次买、第一次卖、第二次买、第二次卖能产生的最大利润.我们分别用 buy1，sell1，buy2，sell2表示
// 那么转移方程式如下
// 		sell2 = buy2 + 本次卖出价
// 		buy2= sell1 - 本次买入价
// 		res[1] = res[0] + 本次卖出价
// 		res[0] = - 本次买入价
func maxProfit3(prices []int) int {
	buy1, sell1, buy2, sell2 := math.MinInt32, 0, math.MinInt32, 0
	for i := 0; i < len(prices); i++ {
		sell2 = Max(sell2, buy2+prices[i])
		buy2 = Max(buy2, sell1-prices[i])
		sell1 = Max(sell1, buy1+prices[i])
		buy1 = Max(buy1, -prices[i])
	}
	return Max(sell1, sell2)
}

func maxProfit4(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = Max(buy1, -prices[i])
		sell1 = Max(sell1, buy1+prices[i])
		buy2 = Max(buy2, sell1-prices[i])
		sell2 = Max(sell2, buy2+prices[i])
	}
	return sell2
}
