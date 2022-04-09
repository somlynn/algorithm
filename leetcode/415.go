package leetcode

import "strconv"

/*
	字符串相加
*/

func addStrings(num1 string, num2 string) string {
	var carry int
	var result string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		sum := a + b + carry
		carry = sum / 10
		result = strconv.Itoa(sum%10) + result
	}
	return result
}
