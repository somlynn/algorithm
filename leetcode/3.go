package leetcode

/*
	无重复字符的最长子串
*/

// 滑动窗口 时间复杂度：O(N),空间复杂度：O(∣Σ∣) ∣Σ∣ = 128，map最多存储128个字符

// 从左向右遍历字符串，用idxMap记录字符最新的索引位置。
// left,right 表示滑动窗口的左右边界，maxLen 记录最长子串的长度。
// 如果重复的元素在滑动窗口内，即idx >= left，则更新左边界 left = idx+1
// 此时比较子串的长度和maxLen,更新maxLen.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, maxLen := 0, 0
	idxMap := make(map[int32]int)
	for right, c := range s {
		if idx, ok := idxMap[c]; ok {
			if idx >= left {
				left = idx + 1
			}
		}
		idxMap[c] = right
		maxLen = Max(right-left+1, maxLen)
	}
	return maxLen
}

// 变形，输出最长子串
func lengthOfLongestSubstring2(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, start, end := 0, 0, 0
	idxMap := make(map[int32]int)
	for right, c := range s {
		if idx, ok := idxMap[c]; ok {
			if idx >= left {
				left = idx + 1
			}
		}
		idxMap[c] = right
		if right-left+1 > end-start {
			start = left
			end = right
		}
	}
	return s[start : end+1]
}
