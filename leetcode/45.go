package leetcode

/*
	跳跃游戏II
	使用最少的跳跃次数到达数组的最后一个位置。
*/

// 贪心,每次尽量跳跃最大位置，就能得到最小次数。
// step 记录步数，即跳跃次数。maxPosition记录能够跳跃的最大位置。

// 假如第一个位置的值为 x ,那么后面的x-1个位置都可以作为下一个起跳点，可以对每一个能作为起跳点的位置都尝试跳一次
// 能跳的最远的距离不断更新。
// 以第一个位置成为第一个起跳点，(1,x] 存在第二个起到点，然后 (x,y]存在第三个起跳点，这些区间的个数就是跳跃次数。
func jump(nums []int) int {
	end, maxPosition, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = Max(maxPosition, i+nums[i])
		// 当i==end 每一个区间遍历完之后，更新次数。
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}
