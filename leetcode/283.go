package leetcode

// 双指针，left 指向零的位置，right找到非零数，然后将right上的数字覆盖left上的数字，当right遍历完了之后，left之后都可以置为零
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			if right != left {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
		right++
	}
}
