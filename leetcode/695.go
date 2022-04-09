package leetcode

/*
	岛屿的最大面积
*/

func maxAreaOfIsland(grid [][]int) int {
	maxArea, r, c := 0, len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 || i >= r || j < 0 || j >= c || grid[i][j] == '0' {
			return 0
		}
		grid[i][j] = '0'
		sum := 1
		sum += dfs(i-1, j)
		sum += dfs(i+1, j)
		sum += dfs(i, j-1)
		sum += dfs(i, j+1)
		return sum
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				maxArea = Max(maxArea, dfs(i, j))
			}
		}
	}
	return maxArea
}
