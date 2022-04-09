package leetcode

import "sort"

/*
	三数之和
*/

// 排序+双指针,时间复杂度: O(N^2) 空间复杂度: O(logN) 额外排序的空间复杂度为O(logN)
func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	n := len(nums)
	ret := make([][]int, 0, 0)
	// 先固定第一个元素，从[0,n-3]遍历，然后利用双指针left，right遍历
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, n-1; left < right; {
			if sum := nums[i] + nums[left] + nums[right]; sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
