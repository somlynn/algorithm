package other

import "strconv"

// 字符串相减(大数相减)

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的差。

func towNumSub(a, b string) string {
	res := ""
	if cmp(a, b) {
		return "-" + towNumSub(b, a)
	}
	i, j := len(a)-1, len(b)-1
	borrow := 0
	for i >= 0 || j >= 0 || borrow > 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(a[i] - '0')
		}
		if j >= 0 {
			y = int(b[j] - '0')
		}
		sub := x - borrow - y
		diff := sub % 10
		borrow = 0
		if sub < 0 {
			diff = (sub + 10) % 10
			borrow = 1
		}
		res = strconv.Itoa(diff) + res
		i--
		j--
	}
	// 清除前导0，比如 104 - 100 = 004
	start := -1
	for i := 0; i < len(res); i++ {
		if res[i] == '0' {
			start = i
		}
	}
	return res[start+1:]
}

func cmp(a, b string) bool {
	if len(a) == len(b) {
		return a <= b
	}
	return len(a) <= len(b)
}
