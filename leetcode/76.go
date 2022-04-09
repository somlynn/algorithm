package leetcode

/*
	最小覆盖子串
*/

// 滑动窗口
// 总体思路：不断扩展滑动窗口右边界，当覆盖所有t时，在尝试收缩左边界，得到结果集，然后从这些结果集中取最小的
// 时间复杂度O(N),空间复杂度O(n)
func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	// 如果len(s) < len(t) 返回空串
	if n < m {
		return ""
	}
	// needMap记录 t 所需要的字符和个数
	needMap := make(map[byte]int)
	for i := range t {
		needMap[t[i]]++
	}
	// [left,right]表示滑动窗口的左右边界，start,end 记录最小子串的开始位置和结束位置
	// needCnt 还有多少个字符等待覆盖。
	left, start, end, needCnt := 0, 0, n, m
	for right := 0; right < n; right++ {
		c := s[right]
		// 如果覆盖了一个字符，就needCnt--
		if needMap[c] > 0 {
			needCnt--
		}
		needMap[c]--
		// 当前needCnt =0 说明已经全覆盖了，但是可能包含多余的元素，开始收缩左边界
		if needCnt == 0 {
			for {
				c := s[left]
				// 当needMap[c] ==0 说明 [left,right] 是以right为右边阶的所有子串最小的了，没有多余的元素。此时得到结果集
				if needMap[c] == 0 {
					break
				}
				// 收缩左边界，去掉多余的元素
				needMap[c]++
				left++
			}
			// 记录最小结果
			if right-left < end-start {
				start = left
				end = right
			}
		}
	}
	// 如果 start，end 从来没有被更新，则不存在结果 返回空串
	if start == 0 && end == n {
		return ""
	}
	return s[start : end+1]
}
