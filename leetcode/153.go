package leetcode

/*
	寻找旋转排序数组中的最小值
*/

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[m] <= nums[h] {
			h = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
