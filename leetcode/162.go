package leetcode

/*
	寻找峰值
*/

// 二分查找
// 二分查找 如果 nums[m] >= nums[m+1] 说明中位数在递减的线段中，因此峰值在[l,m]区间中
// 如果nums[m] < nums[m+1] 说明中位数在递增的线段中，因此峰值在[m+1,h]区间中
func findPeakElement(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[m] >= nums[m+1] {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}
