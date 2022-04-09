package leetcode

/*
	在排序数组中查找元素的第一个和最后一个位置
*/

// 二分查找
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := binarySearchRange(nums, target)
	if left > len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := binarySearchRange(nums, target+1)

	return []int{left, right - 1}
}

// 如果想求左边界的位置，则右边界一定等于中间数 即h=m,让左边界不断收缩，直到l==h
// 如果想求右边界的位置，则左边界一定等于中间数 即l=m,让右边界不断收缩，直到l==h
// 如果等于，则 l 就是 target的左边界
// 如果不存在,则 l 可能是 target-1的右边界的+1,此时 return值可能大于len(nums)，

// 如果不存在 target 左边界，则元素不存在。有两种可能，1、 left 大于len(nums)，2、nums[left]!=target。
func binarySearchRange(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[m] >= target {
			h = m
		} else {
			l = m + 1
		}
	}
	if nums[l] == target-1 {
		return l + 1
	}
	return l
}
