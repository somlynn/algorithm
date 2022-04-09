package leetcode

/*
	编辑距离
*/

// 动态规划
// dp[i][j] 代表 word1 到 i 位置转换成 word2 到 j 位置需要最少步数
// 如果 word1[i] == word2[j] 不需要任何操作，那么dp[i][j] = dp[i-1][j-1]
// 如果 word[i] != word[j] word1 到 i 位置转换成 word2 到 j 位置有如下可能
// 		假如 word1[0,i-1] 与 word2[0,j-1] 匹配了，那么从word1到word2替换即可 dp[i][j]=dp[i-1][j-1]+1
// 		假如word1[0,i-1] 与 word2[0,j] 匹配了，那么从word1到word2 删除word1[i]即可 dp[i][j]=dp[i-1][j]+1
// 		假如word1[0,i] 与 word2[0,j-1] 匹配了，那么从word1到word2  向word1插入word2[j]字符即可 dp[i][j]=dp[i][j-1]+1
// 		因此 dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。
// 第一行第一列，当word1为空或者word2为空时，相互转化就只有删除和增加。
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// dp[0][0] 表示""到""的转换，所以dp[0][0] = 0
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果word1[i] == word2[j],因为下标从0开始。所以word1[i-1] == word2[j-1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(Min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
