package leetcode

/*
	最大子序和
*/

// 动态规划
// dp[i] 表示[0~i]的最大子序和,maxSum记录最大子序和
// if dp[i-1] < 0 dp[i]=nums[i]
// if dp[i-1] > 0 dp[i]= dp[i-1]+nums[i]
// dp[i]再与maxSum 比较，更新maxSum的值
// 时间复杂度：O(N) 空间复杂度： O(1)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxSum = Max(maxSum, nums[i])
	}
	return maxSum
}
