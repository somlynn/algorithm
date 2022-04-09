package other

/*
	双栈排序
*/
func stackSort(stack []int) []int {
	if len(stack) == 0 {
		return nil
	}
	temp := make([]int, 0)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for len(temp) > 0 && temp[len(temp)-1] > top {
			stack = append(stack, temp[len(temp)-1])
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, top)
	}
	return temp
}
