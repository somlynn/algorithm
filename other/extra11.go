package other

// 给定长度为n的数组，每个元素代表一个木头的长度，木头可以任意截断，
// 从这堆木头中截出至少k个相同长度为m的木块。已知k，求max(m)。

// 暴力法
// 从取最长的木头，然后从1~n截取尝试，当满足条件时，取可以截取的最长的
func cutWood(nums []int, k int) int {
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxLen {
			maxLen = nums[i]
		}
	}
	res := 0
	for m := 1; m <= maxLen; m++ {
		cnt := 0
		for i := 0; i < len(nums); i++ {
			cnt += nums[i] / m
		}
		if cnt >= k {
			res = Max(res, m)
		}
	}
	return res
}

// 二分法
func cutWood2(nums []int, k int) int {
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxLen {
			maxLen = nums[i]
		}
	}
	l, h := 0, maxLen
	for l < h {
		m := l + (h-l)/2
		cnt := 0
		for _, num := range nums {
			cnt += num / m
		}
		if cnt >= k {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return l
}
