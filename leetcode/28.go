package leetcode

/*
	实现 strStr()
*/

func strStr(haystack string, needle string) int {
	size1, size2 := len(haystack), len(needle)
	for i := 0; i <= size1-size2; i++ {
		a, b := i, 0
		for b < size2 && haystack[a] == needle[b] {
			a++
			b++
		}
		if b == size2 {
			return i
		}
	}
	return -1
}

// kmp 算法
func strStr2(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	p := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = p[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		p[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = p[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
