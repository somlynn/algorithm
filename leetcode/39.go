package leetcode

import "sort"

/*
	组合总和
*/

// 回溯+剪枝 无重复的candidates 和一个目标数 target
// candidates 中的数字可以无限制重复被选取。

// 快速剪枝：首先排序，如果 target - candidates[i] < 0 ，则后面的直接过滤
// 1、排序,排序的目的是为了剪枝。
// 2、当前target == 0 即找到结果，加入到结果集中
// 3、begin记录开始遍历的位置，如果满足则加入到path中，
//   然后继续从begin开始尝试即backtrack(i,target-candidates[i])，因为数字可以被多次选取。
//  当前以begin作为开始位置遍历完后，path回到最开始的位置，在进行一下元素的遍历，即回溯。
func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	var backtrack func(int, int)
	backtrack = func(begin int, target int) {
		if target == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ret = append(ret, copyPath)
			return
		}
		for i := begin; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return ret
}
