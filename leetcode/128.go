package leetcode

/*
	最长连续序列
*/

// 哈希表
func longestConsecutive(nums []int) int {
	maxLen := 0
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	for num := range numMap {
		if !numMap[num-1] {
			cur := num
			for numMap[cur] {
				cur++
			}
			maxLen = Max(maxLen, cur-num)
		}
	}
	return maxLen
}
