package leetcode

/*
	只出现一次的数字
*/

// 位运算: 异或
// a ^ 0 = a
// a ^ a = 0
// 异或运算满足交换律和结合律：a^b^a = (a^a)^b=0^b=b
func singleNumber(nums []int) int {
	var m int
	for _, v := range nums {
		m ^= v
	}
	return m
}
