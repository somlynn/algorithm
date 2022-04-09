package leetcode

/*
	单词搜索
*/

func exist(board [][]byte, word string) bool {
	r, c := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i int, j int, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= r || j < 0 || j >= c {
			return false
		}

		if word[k] != board[i][j] {
			return false
		}
		t := board[i][j]
		board[i][j] = '0'
		res := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = t
		return res
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
