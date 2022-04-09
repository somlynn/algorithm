package leetcode

import "sort"

/*
	多数元素
*/

// 多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

// 哈希表
func majorityElement(nums []int) int {
	countMap := make(map[int]int, len(nums))
	for _, v := range nums {
		countMap[v]++
	}
	for k, v := range countMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

// 排序
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Boyer-Moore(摩尔)投票算法
func majorityElement4(nums []int) int {
	candidate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
