package leetcode

import "math"

/*
	乘积最大子数组
*/
func maxProduct(nums []int) int {
	// max,min 用来记录nums[0,i-1]的最大值和最小值
	res, max, min := math.MinInt32, 1, 1
	for i := 0; i < len(nums); i++ {
		// 当nums[i] < 0 是 nums[i]与max，min乘积是相反的
		if nums[i] < 0 {
			max, min = min, max
		}
		max = Max(nums[i], max*nums[i])
		min = Max(nums[i], min*nums[i])
		res = Max(res, max)
	}
	return res
}
