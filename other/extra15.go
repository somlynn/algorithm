package other

import (
	"math"
	"sort"
)

// 两个数组的最小差

// 排序+双指针
// 首先对两个数组进行从小到大排序,
// 如果 a[i] < b[j] ,如果b[j]-a[i] 小于最小值，则更新最小值,
// 此时如果，移动 j a[i],b[j]的差值会更大，所以我们移动i尝试更新。
// 如此 重复操作。知道一个数组遍历完毕。
func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	i, j := 0, 0
	minDiff := math.MaxInt32
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			minDiff = Min(minDiff, b[j]-a[i])
			i++
		} else {
			minDiff = Min(minDiff, a[i]-b[j])
			j++
		}
	}
	return minDiff
}
