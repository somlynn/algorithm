package leetcode

/*
	爬楼梯
*/

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	p1, p2, cur := 1, 2, 0
	for i := 3; i <= n; i++ {
		cur = p1 + p2
		p1 = p2
		p2 = cur
	}
	return cur
}
