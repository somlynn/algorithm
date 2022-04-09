package leetcode

import (
	"fmt"
	"sort"
)

// 组合总和 II

//  数组中的每个数字在每个组合中只能使用一次。
// 此题和39的不同点是，每个元素只能使用一次，并且结果集不能有重复。
// begin 表示 以begin位置作为起点，然后向下尝试，遍历完了之后，我们要回退，然以begin+1作为起点再尝试。
// 以 nums= 10,1,2,7,6,1,5  target=8 为例
// 先排序，1,1,2,5,6,7,10
// [1]
// [1,1]
// [1,1,2]
// [1,1,5]
// [1,1,6] -> res
// [1,2,5]
// [1,5,6]
// [1,7]
//
func combinationSum2(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)
	var backtrack func(int, int)
	backtrack = func(begin int, target int) {
		if target == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ret = append(ret, copyPath)
		}
		for i := begin; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			fmt.Println("path=", path)
			backtrack(i+1, target-candidates[i])
			path = path[:len(path)-1]
			fmt.Println("path=", path)
		}
	}
	backtrack(0, target)
	return ret
}
