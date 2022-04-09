package leetcode

/*
	打家劫舍
*/

// 动态规划 dp[i] 表示打劫[0,i]范围内住宅的最大价值，
// 由于相邻的不能打劫，dp[i] = max(dp[i-1],dp[i-2]+nums[i])
// 但是用 dp数组来解决问题，会带来额外的空间,
// 从上述公式看，dp[i]只和 dp[i-1] dp[i-2]有关，我们可以使用变量代替
func rob(nums []int) int {
	cur, pre, preNext := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		cur = Max(pre, preNext+nums[i])
		preNext = pre
		pre = cur
	}
	return pre
}
