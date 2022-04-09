package leetcode

import "math"

/*
	长度最小的子数组
*/

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen := math.MaxInt32
	start, end, sum := 0, 0, 0
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			maxLen = Min(maxLen, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if maxLen == math.MinInt32 {
		return 0
	}
	return maxLen
}
