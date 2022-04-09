package leetcode

import "math"

/*
	整数反转
*/

func reverse(x int) int {
	var res int
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
	return res
}
