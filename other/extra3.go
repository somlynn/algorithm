package other

/*
	36进制由0-9，a-z，共36个字符表示。

	要求按照加法规则计算出任意两个36进制正整数的和，如1b + 2x = 48  （解释：47+105=152）

	要求：不允许使用先将36进制数字整体转为10进制，相加后再转回为36进制的做法
*/

// 这个题目其实是415 两个字符串相加的扩展

func add36Strings(num1 string, num2 string) string {
	carry := 0
	var result string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = getInt(num1[i])
		}
		if j >= 0 {
			b = getInt(num2[i])
		}
		sum := a + b + carry
		carry = sum / 36
		result = string(getChar(sum%36)) + result
	}
	return result
}

func getInt(ch byte) int {
	if ch >= '0' && ch < '9' {
		return int(ch - '0')
	} else {
		return int(ch-'a') + 10
	}
}

func getChar(num int) byte {
	if num >= 0 && num < 9 {
		return '0' + byte(num)
	} else {
		return byte(num-10) + 'a'
	}
}
