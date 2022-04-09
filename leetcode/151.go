package leetcode

import "strings"

/*
	翻转字符串里的单词
*/
// 首先我们找到第一个空格，那么[i+1,j]组成了一个单词。这里注意一下，字符串首尾可能存在空格，
// 所以[i+1,j] 不等于' '才加入结果中
// 然后滤掉空格，找到第一个不为空的字符。然后更新j=i，使j指向下一个单词的尾部。
func reverseWords(s string) string {
	ret := make([]string, 0)
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for ; i >= 0 && s[i] != ' '; i-- {
		}
		if i <= j && s[j] != ' ' {
			ret = append(ret, s[i+1:j+1])
		}
		for ; i >= 0 && s[i] == ' '; i-- {
		}
		j = i
	}
	return strings.Join(ret, " ")
}
