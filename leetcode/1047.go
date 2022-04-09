package leetcode

// 使用栈
func removeDuplicates2(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
