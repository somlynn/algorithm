package leetcode

import "strings"

/*
	验证回文串
*/

func isPalindrome3(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isLetterOrDigit(s[i]) {
			i++
		}
		for i < j && !isLetterOrDigit(s[j]) {
			j--
		}
		if i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func isLetterOrDigit(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
