package leetcode

/*
	最长公共子序列
*/

// 动态规划
// 总体思想：dp[i][j]表示 text1[0~i] 与 text2[0~j] 最长公共子序列
// if text1[i] == text2[j] 那么dp[i][j] = dp[i-1][j-1] +1
// if text1[i] != text2[j] 可能为 text1[0,i-1],text2[0,j]的最长公共子序列 或者 text1[0,i],text2[0,j-1]。我们取最大的
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果text1[i] == text2[j] ,那么在原来的最长公共子序列dp[i-1][j-1]+1即可
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 如果text1[i] != text2[j],text1[0~i] 与 text2[0~j]的最长公共子序列,
				// 可能为 text1[0~i-1],text2[0~j]的最长公共子序列 或者 text1[0~i],text2[0~j-1]。我们取最大的
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
