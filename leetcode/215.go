package leetcode

/*
	数组中的第K个最大元素
	topK 问题
*/

// 因为是求第K个最大元素，position 切分之后应该是 左边大于切分元素，右边小于切分元素。如果求第K小个数，则相反。
// 第k个大元素，则对应的切分元素的下标应该是k-1.
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return -1
	}
	l, h := 0, len(nums)-1
	for l < h {
		j := position(nums, l, h)
		if j == k-1 {
			break
		} else if j < k-1 {
			l = j + 1
		} else {
			h = j - 1
		}
	}
	return nums[k-1]
}

// 拿第一个元素作为切分元素v，i,j 双指针与v比较
// i直到找到第一个大于v的数，j 直到找到第一个小于v的数
// 交换i,j的值。重复进行直到i>=j为止，最后交换l,j的值

func position(nums []int, l, h int) int {
	v := nums[l]
	i, j := l+1, h
	for {
		for ; i < h && nums[i] >= v; i++ {
		}
		for ; j > l && nums[j] <= v; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
