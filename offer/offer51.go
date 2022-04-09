package offer

/*
	数组中的逆序对
*/

var aux []int

func reversePairs(nums []int) int {
	aux = make([]int, len(nums))
	return mergeSort2(nums, 0, len(nums)-1)
}

func mergeSort2(nums []int, l, h int) int {
	if l >= h {
		return 0
	}
	m := l + (h-l)/2
	left := mergeSort2(nums, l, m)
	right := mergeSort2(nums, m+1, h)
	return mergeList2(nums, l, m, h) + left + right
}

func mergeList2(nums []int, l, m, h int) int {
	if l >= h {
		return 0
	}
	i, j := l, m+1
	count := 0
	for k := l; k <= h; k++ {
		aux[k] = nums[k]
	}
	for k := l; k <= h; k++ {
		if i > m {
			nums[k] = aux[j]
			j++
		} else if j > h {
			nums[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			count += m - i + 1
		}
	}
	return count
}
