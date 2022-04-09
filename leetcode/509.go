package leetcode

func fib(n int) int {
	if n < 2 {
		return n
	}
	p1, p2, cur := 0, 1, 0
	for i := 2; i <= n; i++ {
		cur = p1 + p2
		p1 = p2
		p2 = cur
	}
	return cur
}
