package leetcode

/*
	x 的平方根
*/
// 比如以15为例，8
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 1, x
	for l <= h {
		m := l + (h-l)/2
		sqrt := x / m
		if sqrt == m {
			return sqrt
		} else if sqrt < m {
			h = m - 1
			l = sqrt + 1
		} else {
			return m
		}
	}
	return h
}

// 变形
func mySqrt2(x float64, precision float64) float64 {
	if x == 0 {
		return 0
	}
	l, h := 0.0, x
	for l <= h {
		m := l + (h-l)/2.0
		if m*m-x <= precision {
			return m
		} else if m*m-x > precision {
			h = m
		} else {
			l = m
		}
	}
	return -1
}
