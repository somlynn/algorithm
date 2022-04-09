package leetcode

import "fmt"

/*
	子集
*/

// 迭代
func subsets(nums []int) [][]int {
	res := make([][]int, 0, 0)
	size := len(nums)
	for mask := 0; mask < 1<<size; mask++ {
		set := make([]int, 0, 0)
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		res = append(res, set)
	}
	return res
}

// 递归(回溯)
// cur表示当前位置，原数组的每个位置都有两种状态，取和不取。我们用set存放被选中的数字
// backtrack(cur)之前说明[0,cur-1]是确定，[cur,len(nums))是还没确定的。
// 如果取nums[cur]，进行backtrack(cur+1),遍历完了之后，回溯set = set[:len(set)-1]
// 此时set集合不含有nums[cur],然后我们不取nums[cur] 直接进行一下 backtrack(cur+1)

func subsets2(nums []int) [][]int {
	res := make([][]int, 0)
	set := make([]int, 0)
	var backtrack func(int)
	backtrack = func(cur int) {
		if cur == len(nums) {
			cp := make([]int, len(set))
			copy(cp, set)
			fmt.Println(cp)
			res = append(res, cp)
			return
		}
		// 取nums[cur]
		set = append(set, nums[cur])
		fmt.Println(set)
		backtrack(cur + 1)
		// 回退了，此时set不含nums[cur]
		set = set[:len(set)-1]
		// 不取nums[cur],不添加到set，直接backtrack
		backtrack(cur + 1)
	}
	backtrack(0)
	return res
}
