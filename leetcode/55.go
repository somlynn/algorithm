package leetcode

// 跳跃游戏

// 贪心
// 从左往右遍历，用maxJump记录能够到达的最大距离的下标。
// 如果能到达i,j,k 位置的话，则 maxJump = Max(i+nums[i],j+nums[j],k+nums[k])
// 所以如果 i<= maxJump 说明能够到达i，那i+nums[i]就是从i开始能到达的最大位置。在和之前的元素的最大位置比较，
// 更新maxJump.if maxJump >= len(nums)-1说明任何节点都可以到达则直接返回true。
func canJump(nums []int) bool {
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxJump {
			maxJump = Max(maxJump, i+nums[i])
			if maxJump >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
