package leetcode

/*
	最长回文子串
*/

// 中心扩展法
// 时间复杂度 O(N^2) 空间复杂度(1)
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		s1, e1 := extendCenter(s, i, i)
		s2, e2 := extendCenter(s, i, i+1)
		if e1-s1 > end-start {
			start, end = s1, e1
		}
		if e2-s2 > end-start {
			start, end = s2, e2
		}
	}
	return s[start : end+1]
}

func extendCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
