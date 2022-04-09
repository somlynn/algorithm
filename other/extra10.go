package other

// 将正数移动到数组左边，负数移动到数组右边,并且相对顺序不变

func moveNums(nums []int) []int {
	queue := make([]int, 0)
	queue2 := make([]int, 0)
	for _, v := range nums {
		if v < 0 {
			queue = append(queue, v)
		} else {
			queue2 = append(queue2, v)
		}
	}
	result := make([]int, 0)
	result = append(result, queue...)
	result = append(result, queue2...)
	return result
}
