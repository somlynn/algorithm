package leetcode

import "sort"

/**
全排列II
*/

// 回溯+剪枝
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	perm := []int{}
	visited := make([]bool, size)
	res := make([][]int, 0, 0)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == size {
			res = append(res, append([]int{}, perm...))
			return
		}
		for i, v := range nums {
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			visited[i] = true
			dfs(idx + 1)
			visited[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	dfs(0)
	return res
}
