package leetcode

// 组合
// 回溯
func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	var backtrack func(int, int)
	backtrack = func(begin int, k int) {
		if len(path) == k {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ret = append(ret, copyPath)
		}
		for i := begin; i <= n; i++ {
			path = append(path, i)
			backtrack(i+1, k)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, k)
	return ret
}
