package leetcode

/*
	括号生成
*/

// 回溯
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var backtrack func(string, int, int)
	backtrack = func(cur string, open int, close int) {
		if len(cur) == n*2 {
			res = append(res, cur)
		}
		if open < n {
			backtrack(cur+"(", open+1, close)
		}
		if close < open {
			backtrack(cur+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return res
}
