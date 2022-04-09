package leetcode

import "sort"

// 排序
// 时间复杂度为 O(NlogN)，空间复杂度为O(logN)
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

// 哈希表
func missingNumber2(nums []int) int {
	numMap := make(map[int]bool)
	for i := range nums {
		numMap[nums[i]] = true
	}
	for i := 0; i <= len(nums); i++ {
		if !numMap[i] {
			return i
		}
	}
	return len(nums)
}

// 位运算 x^x = 0  0^x = x
func missingNumber3(nums []int) (xor int) {
	// 展开就是  1^1^2^2^3^3....xor...(n-1)^(n-1)^n = xor^n
	for i, num := range nums {
		xor ^= i ^ num
	}
	// xor^n^n = xor^0=xor
	return xor ^ len(nums)
}

// 数学 根据数学公式[0,n] 的和为 total = n*(n+1)/2 ,target = total - sum(nums)
func missingNumber4(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}
