package leetcode

/*
	字典序的第K小数字
*/

// 什么是字典序？
// 根据数字的前缀进行排序，比如10 < 9.112 < 12
// 如果root为1，孩子节点为0~9，那么整个字典序排列也就是对十叉树进行先序遍历

func findKthNumber(n int, k int) int {
	// cur作为指针，指向当前位置，当cur==k时也就是到了第k个数。prefix表示前缀
	cur, prefix := 1, 1
	for cur < k {
		count := getCount(prefix, n)
		// 说明第k个数在该前缀下
		if cur+count > k {
			prefix *= 10
			cur++
		} else {
			prefix++
			cur += count
		}
	}
	return prefix
}

// 获取以prefix为前缀，以n为上界所有元素的个数
func getCount(prefix int, n int) int {
	// count记录个数，cur是前缀，next是下一个前缀
	count, cur, next := 0, prefix, prefix+1
	for cur <= n {
		count += Min(next, n+1) - cur
		cur *= 10
		next *= 10
	}
	return count
}
