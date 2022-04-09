package leetcode

import "strings"

/*
	N皇后
*/

// 回溯+剪枝
func solveNQueens(n int) [][]string {
	queues := make([][]string, n)
	for i := range queues {
		queues[i] = make([]string, n)
		for j := range queues[i] {
			queues[i][j] = "."
		}
	}
	res := make([][]string, 0, 0)
	cols, diag1, diag2 := make(map[int]bool), make(map[int]bool), make(map[int]bool)

	var backtrack func([][]string, int, int, map[int]bool, map[int]bool, map[int]bool)
	backtrack = func(queues [][]string, row int, n int, cols map[int]bool, diag1 map[int]bool, diag2 map[int]bool) {
		if row == n {
			temp := make([]string, len(queues))
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(queues[i], "")
			}
			res = append(res, temp)
		}
		for i := 0; i < n; i++ {
			if !cols[i] && !diag1[row+i] && !diag2[row-i] {
				queues[row][i] = "Q"
				cols[i] = true
				diag1[row+i] = true
				diag2[row-i] = true
				backtrack(queues, row+1, n, cols, diag1, diag2)
				queues[row][i] = "."
				cols[i] = false
				diag1[row+i] = false
				diag2[row-i] = false
			}
		}
	}
	backtrack(queues, 0, n, cols, diag1, diag2)
	return res
}
