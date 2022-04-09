package leetcode

/*
	缺失的第一个正数
*/

// 通过置换让元素回归到自己的位置上，比如 2 应该在第二个位置上，即nums[1]==2,
// 置换结束后，遍历数组看哪个元素没有回归到自己的位置上。
// 注意数字与下标的关系

// 时间复杂度为O(N),空间复杂度为O(1)
// 这里解析下为啥是O(N)，不能因为有两个for循环就是O(N^2),里边的for不是每次循环都执行的，
// 第二个for每次循环都会是其中的一个整数归位，如果已经归位好了，直接跳过的，即nums[nums[i]-1] == nums[i]
// 所以整时间复杂度和交换次数有关，交换次数最多为N,而比较次数是N^2。
// 所以时间复杂度为O(N)
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
