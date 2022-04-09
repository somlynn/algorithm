package other

/*
	在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。求一个数组的小和。
	要求时间复杂度O(NlogN)，空间复杂度O(N)
*/

// 归并排序的思想
// 记数组最小和 smallsum
// smallsum(nums[1,3,4,2,5]) = smallsum([1,3,4]) + smallsum([2,5]) + c
// c 是指左半边小于右半边的和，这个和我们可以使用归并排序的技巧

var bet []int

func getSmallSum(nums []int) int {
	bet = make([]int, len(nums))
	return getSum(nums, 0, len(nums)-1)
}

func getSum(nums []int, l, h int) int {
	if l >= h {
		return 0
	}
	m := l + (h-l)/2
	left := getSum(nums, l, m)
	right := getSum(nums, m+1, h)
	return mergeList(nums, l, m, h) + left + right
}

func mergeList(nums []int, low, mid, high int) int {
	sum := 0
	i, j := low, mid+1
	for k := low; k <= high; k++ {
		bet[k] = nums[k]
	}
	for k := low; k <= high; k++ {
		if i > mid {
			nums[k] = bet[j]
			j++
		} else if j > high {
			nums[k] = bet[i]
			i++
		} else if bet[i] < bet[j] {
			nums[k] = bet[i]
			sum += (high - j + 1) * bet[i]
			i++
		} else {
			nums[k] = bet[j]
			j++
		}
	}
	return sum
}
