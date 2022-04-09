package leetcode

/*
	搜索旋转排序数组
*/

// 二分查找
// 左边的元素比较大，右边的元素较小。 中间的元素：nums[m]
// if nums[m] == target,则直接返回
// if if nums[0] <= nums[m] ,则中间的元素在左半部，否则在右半部
// 如果在左半部：
// 		if nums[0] <= target && target < nums[m], 则target一定在[0,m)之间否则在(m,n]
// 如果在右半部：
// 		if nums[m] < target && target <= nums[len(nums)-1] 则target一定在(m,n],否则在[0,m)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[0] <= nums[m] {
			if nums[0] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[len(nums)-1] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
