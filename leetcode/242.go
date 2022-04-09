package leetcode

// 哈希表
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	byteMap := make(map[int32]int, 26)
	for _, v := range s {
		byteMap[v]++
	}
	for _, v := range t {
		byteMap[v]--
		if byteMap[v] < 0 {
			return false
		}
	}
	return true
}
