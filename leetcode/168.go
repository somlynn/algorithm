package leetcode

// 26进制转换
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
// 28/26= 1 -> A 28%26=2 -> B == AB
func convertToTitle(columnNumber int) string {
	ret := make([]byte, 0)
	for columnNumber > 0 {
		// 减一的目的是 为了 'A'+byte(columnNumber%26) ,比如 26 % 26 == 0 ， 'A'+0 = 'A',所以这里减一为了求偏移
		columnNumber--
		ret = append(ret, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	i, j := 0, len(ret)-1
	for i < j {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}
	return string(ret)
}
