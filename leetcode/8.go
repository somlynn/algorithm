package leetcode

import (
	"math"
)

/*
	字符串转换整数 (atoi)
*/

func myAtoi(s string) int {
	i, sign, res := 0, 1, 0
	if len(s) == 0 {
		return 0
	}
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) {
		digit := s[i] - '0'
		if digit < 0 || digit > 9 {
			break
		}
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(digit)
		i++
	}
	return res * sign
}
