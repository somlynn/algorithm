package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool)
	retMap := make(map[int]bool)
	for _, num := range nums1 {
		numMap[num] = true
	}
	for _, num := range nums2 {
		if numMap[num] {
			retMap[num] = true
		}
	}
	ret := make([]int, 0, len(retMap))
	for key := range retMap {
		ret = append(ret, key)
	}
	return ret
}
