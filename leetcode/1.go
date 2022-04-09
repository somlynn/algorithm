package leetcode

/*
	两数之和
*/

// 一次遍历+哈希表
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	numIdxMap := make(map[int]int)
	for i, v := range nums {
		if n, ok := numIdxMap[target-v]; ok {
			return []int{n, i}
		}
		numIdxMap[v] = i
	}
	return nil
}
