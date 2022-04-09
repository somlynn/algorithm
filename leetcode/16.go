package leetcode

import (
	"math"
	"sort"
)

/*
	最接近的三数之和
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if Abs(sum-target) < Abs(res-target) {
				res = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
