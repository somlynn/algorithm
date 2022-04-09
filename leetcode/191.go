package leetcode

// 观察这个运算：n &(n - 1)，其运算结果恰为把 n 的二进制位中的最低位的 1 变为 0之后的结果。
// 这样我们可以利用这个位运算的性质加速我们的检查过程，在实际代码中，
// 我们不断让当前的 n 与 n - 1 做与运算，直到 n 变为 0 即可。
// 因为每次运算会使得 n 的最低位的 1 被翻转，因此运算次数就等于 n 的二进制位中 1 的个数。
func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}
