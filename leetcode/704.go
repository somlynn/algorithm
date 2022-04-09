package leetcode

/*
	二分查找
*/

func search2(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}
