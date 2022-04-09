package leetcode

/*
	Pow(x, n)
*/

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow2(x float64, n int) float64 {
	if n >= 0 {
		return quickMul2(x, n)
	}
	return 1 / quickMul2(x, -n)
}

func quickMul2(x float64, n int) float64 {
	res := 1.0
	y := x
	for n > 0 {
		if n%2 == 1 {
			res *= y
		}
		y *= y
		n /= 2
	}
	return res
}
