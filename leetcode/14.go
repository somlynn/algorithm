package leetcode

/*
	最长公共前缀
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s0 := strs[0]
	for i, c := range s0 {
		for _, s := range strs {
			if i == len(s) || int32(s[i]) != c {
				return s0[:i]
			}
		}
	}
	return s0
}
