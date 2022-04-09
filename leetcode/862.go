package leetcode

/*
	和至少为K的最短子数组
*/

// 滑动窗口
func shortestSubarray(nums []int, k int) int {
	size := len(nums)
	preSum := make([]int, size+1)
	for i := 0; i < size; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	queue := make([]int, 0)
	res := size + 1
	for i := 0; i < len(preSum); i++ {
		for len(queue) != 0 && preSum[i] <= preSum[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		for len(queue) != 0 && preSum[i] >= preSum[queue[0]]+k {
			res = Min(res, i-queue[0])
			queue = queue[1:]
		}
		queue = append(queue, i)
	}
	if res < size+1 {
		return res
	}
	return -1
}
