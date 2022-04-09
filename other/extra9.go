package other

// 求区间最小数乘区间和的最大值

// 给定一个数组，要求选出一个区间, 使得该区间是所有区间中经过如下计算的值最大的一个：区间中的最小数 * 区间所有数的和。

// 暴力法
//题目是找max(区间和 * 区间最小值)，而满足的区间最小值一定是数组的某个元素。
//因此可以枚举数组，枚举的每个元素（设为x）作为区间最小值，在x左右两侧找到第一个比x小的元素，
//分别记录左右边界的下标为l,r，寻找边界时计算当前区间的和。
//那么以x为区间最小值的最大计算区间就是[l+1,r-1]区间和*x，枚举时更新最大值。
//整个算法的时间复杂度是O(N^2),空间复杂度为O(1)
func rangeMaxSum(nums []int) int {
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		left, right := i-1, i+1
		leftSum, rightSum := 0, 0
		for left >= 0 && nums[left] >= nums[i] {
			leftSum += nums[left]
			left--
		}
		for right < len(nums) && nums[right] >= nums[i] {
			rightSum += nums[right]
			right++
		}
		maxSum = Max(maxSum, (leftSum+rightSum+nums[i])*nums[i])
	}
	return maxSum
}

// 优化
// 单调栈。方法一中找每个元素左右边界的复杂度是O(N)，通过单调栈的数据结构可以将其优化为O(1)。
// 因此优化后整个算法的时间复杂度可以达到O(N)。
// 1 2 3 4 5
// 1 3 6 10 15
func rangeMaxSum2(nums []int) int {
	// 这里求和可以使用前缀和，[l,r]的和等于 sum[r]-sum[l-1]
	// 因为l-1可能小于0,方便计算我们可以从下标1开始
	if len(nums) == 0 {
		return 0
	}
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	maxSum := 0
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] <= nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftSum := 0
			if len(stack) > 0 {
				leftSum = sum[stack[len(stack)-1]]
			}
			maxSum = Max(maxSum, (sum[i-1]-leftSum)*nums[top])
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftSum := 0
		if len(stack) > 0 {
			leftSum = sum[stack[len(stack)-1]]
		}
		maxSum = Max(maxSum, (sum[len(nums)-1]-leftSum)*nums[top])
	}
	return maxSum
}
