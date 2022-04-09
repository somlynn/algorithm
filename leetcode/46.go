package leetcode

/*
	全排列
*/

// 回溯+剪枝
func permute(nums []int) [][]int {
	res := make([][]int, 0, 0)
	visited := make(map[int]bool)
	path := make([]int, 0)
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, num := range nums {
			if visited[num] {
				continue
			}
			path = append(path, num)
			visited[num] = true
			backtrack()
			path = path[:len(path)-1]
			visited[num] = false
		}
	}
	backtrack()
	return res
}
