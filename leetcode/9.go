package leetcode

/*
	回文数
*/

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}
	sum := 0
	for x > sum {
		sum = sum*10 + x%10
		x /= 10
	}
	// 如果x位数为偶数 x == sum 如果x位数为奇数 x == sum/10
	return x == sum || x == sum/10
}
