package leetcode

/*
	分发糖果
*/

// 双指针
// 总体思路：从左向右分配糖果（只比较左边的孩子），得出每位孩子所需要的糖果left[i]。然后从右向左分配糖果（只比较右边的孩子），得出每位孩子所需要的糖果right[i]。
// 因为每位孩子要满足两边孩子的条件，所以 num[i] = Max(left[i],right[i])
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	left := make([]int, len(ratings))
	for i := range ratings {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right, ret := 0, 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ret += Max(left[i], right)
	}
	return ret
}
