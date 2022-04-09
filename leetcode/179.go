package leetcode

import (
	"sort"
	"strconv"
)

/*
	最大数
*/

// 排序
// 思路： 对于 numsnums 中的任意两个值 aa 和 bb，我们无法直接从常规角度上确定其大小/先后关系。
// 但我们可以根据「结果」来决定 aa 和 bb 的排序关系：
// 如果拼接结果 abab 要比 baba 好，那么我们会认为 aa 应该放在 bb 前面。
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		// 先要判断 a，b的位数然后才好拼接
		bit1, bit2 := 10, 10
		for a >= bit1 {
			bit1 *= 10
		}
		for b >= bit2 {
			bit2 *= 10
		}
		return bit2*a+b > bit1*b+a
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := make([]byte, 0)
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
