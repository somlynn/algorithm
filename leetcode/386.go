package leetcode

/*
	字典序排数
*/

// 递归
func lexicalOrder(n int) []int {
	ret := []int{}
	for i := 1; i <= 9; i++ {
		ret = append(ret, dfs(i, n)...)
	}
	return ret
}

func dfs(i, n int) []int {
	ret := []int{}
	if i > n {
		return []int{}
	}
	ret = append(ret, i)
	for j := 0; j <= 9; j++ {
		ret = append(ret, dfs(i*10+j, n)...)
	}
	return ret
}

// 前缀树
func lexicalOrder2(n int) []int {
	ret := make([]int, 0)
	num := 1
	for {
		if num <= n {
			ret = append(ret, num)
			num *= 10
		} else {
			num /= 10
			for num%10 == 9 {
				num /= 10
			}
			if num == 0 {
				break
			}
			num++
		}
	}
	return ret
}
