package leetcode

/*
	字符串解码
*/

// 本题难点在于括号内嵌套括号，需要从内向外生成与拼接字符串
// 使用辅助栈
func decodeString(s string) string {
	// str 存储每个[]内的字符串,存储括号外的倍数
	rate := 0
	str := make([]byte, 0)

	// 使用两个栈，一个存倍数d，一个存字符串str
	strStack := make([]string, 0)
	rateStack := make([]int, 0)
	for i := range s {
		if s[i] == '[' {
			rateStack = append(rateStack, rate)
			strStack = append(strStack, string(str))
			rate = 0
			str = make([]byte, 0)
		} else if s[i] == ']' {
			d := rateStack[len(rateStack)-1]
			rateStack = rateStack[:len(rateStack)-1]
			tempStr := ""
			for j := 0; j < d; j++ {
				tempStr += string(str)
			}

			str = []byte(strStack[len(strStack)-1] + tempStr)
			strStack = strStack[:len(strStack)-1]
		} else if s[i] >= '0' && s[i] <= '9' {
			rate = rate*10 + int(s[i]-'0')
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
