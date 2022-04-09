package leetcode

/*
	字符串相乘
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		a := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			b := num2[j] - '0'
			sum := res[i+j+1] + int(a*b)
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	if len(res) > 0 && res[0] == 0 {
		res = res[1:]
	}
	s := make([]byte, 0)
	for i := range res {
		s = append(s, byte(res[i]+'0'))
	}
	return string(s)
}
