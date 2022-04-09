package leetcode

import "strings"

/*
	移掉K位数字,使剩下的数字最小
*/

// 贪心+单调栈
// 要保证数字最小，我们要保证前面位的数字最小。尽量使小的数字放在前面。
// 因此我们可以使用单调递增栈，
func removeKDigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	// 数组遍历完了，移除的元素还不够K个数，然后移掉剩下的K个数。
	stack = stack[:len(stack)-k]
	// 去除前导0
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
