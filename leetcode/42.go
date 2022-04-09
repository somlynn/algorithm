package leetcode

/*
	接雨水
*/

// 暴力法
// 每一格所能接到的雨水 = if （左右两边高度的最值 > 当前高度）则为高度之差，否则盛不了水，为0，
func trap1(height []int) int {
	var sum int
	for i := range height {
		maxLeft, maxRight := 0, 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		for j := i + 1; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		min := Min(maxLeft, maxRight)
		if min > height[i] {
			sum += min - height[i]
		}
	}
	return sum
}

// 暴力法每次找最大值，都需要向左向右遍历一次，我们可以存储计算的值，减少遍历次数

// 动态规划
// 核心思路：按列求第i位置的积水 = Min(左边最大高度,右边最大高度) - height[i]，注意求左右两边最大值是包含自身height[i]的
// 时间复杂度: O(N),空间复杂度O(N)
func trap2(height []int) int {
	// 当只有两个柱子时，形成不了积水
	if len(height) <= 2 {
		return 0
	}
	var sum int
	size := len(height)
	left := make([]int, size)
	right := make([]int, size)
	left[0] = height[0]
	right[size-1] = height[size-1]
	// 从左向右遍历数组，记录[0,i]的最大值
	for i := 1; i < size; i++ {
		left[i] = Max(height[i], left[i-1])
	}
	// 从右向左遍历数组,记录[i,n]的最大值
	for i := size - 2; i >= 0; i-- {
		right[i] = Max(height[i], right[i+1])
	}
	// i,size-1 位置形成不了积水
	for i := 1; i < size-1; i++ {
		sum += Min(left[i], right[i]) - height[i]
	}
	return sum
}

// 使用单调递减栈，栈存储的是每一个元素的下标。
// 核心思路：
//		1、遍历数组height[i]和栈顶比较，如果大于栈顶元素，此时 height[i]、栈顶元素topItem、栈顶的下一个元素nextItem形成了积水。
// 			sum = min(height[i],height[nextIdx]) * (i-nextIdx-1)、
//		2、如果小于等于栈顶元素，加入栈中
//
// 时间复杂度: O(N),空间复杂度O(N)
func trap3(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	var sum, i int
	size := len(height)
	stack := make([]int, 0, size)
	for i < size {
		// 栈不为空且height[i]大于栈顶元素
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			// 弹出栈顶元素
			topItem := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			// 此时不存在栈顶的下一个元素，形成不了积水。直接返回
			if len(stack) == 0 {
				break
			}
			nextIdx := stack[len(stack)-1]
			distance := i - nextIdx - 1
			min := Min(height[i], height[nextIdx])
			sum += (min - topItem) * distance
		}
		stack = append(stack, i)
		i++
	}
	return sum
}

// 双指针法
// 从左右两边同时计算，如果左边小于右边，则记录左边的最大值，并求和。相反，同理。
// 定理一：在某个位置i处，它能存的水，取决于它左右两边的最大值中较小的一个。
// 定理二：当我们从左往右处理到left下标时，左边的最大值left_max对它而言是可信的， 但right_max对它而言是不可信的。
// 定理三：当我们从右往左处理到right下标时，右边的最大值right_max对它而言是可信的，但left_max对它而言是不可信的。

// 对于left位置，他的左边最大值一定是left_max，右边的最大值 >= right_max。
// 对于right位置，他的左边最大值一定是right_max,左边的最大值 >= left_max。

// 如果 left_max <= right_max <= 右边的最大值.左右两边的最大值最小的一定是left_max.
//			那left上的积水一定是 left_max-height[left]
// 如果 right_max <= left_max <= 左边的最大值.左右两边的最大值最小的一定是right_max.
//			那right上的积水一定是 right_max-height[right]
// 时间复杂度: O(N),空间复杂度O(1)
func trap4(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	var sum int
	maxLeft, maxRight := height[0], height[len(height)-1]
	i, j := 1, len(height)-2
	for i <= j {
		if maxLeft < maxRight {
			maxLeft = Max(maxLeft, height[i])
			sum += maxLeft - height[i]
			i++
		} else {
			maxRight = Max(maxRight, height[j])
			sum += maxRight - height[j]
			j--
		}
	}
	return sum
}
