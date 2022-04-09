package leetcode

/*
	下一个排列
*/

// 13598742 从右向左找到一组递增序列135 98742 此时nums[i] = 5
// 然后从右向左找到第一个大于nums[i]的数，即7 交换127 98542 此时后面的序列依然是递增的，
// nums[i]此时 变成了比5大的最小的数，然后反转后面的递增序列得到 13724589即为下一排列
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse2(nums, i+1)
}

func reverse2(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
