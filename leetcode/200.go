package leetcode

/*
	岛屿数量
*/

// 深度优先
// dfs 主要是用来标记，遇到第一个1，然后dfs把相连的所有1都标记0，然后计为一个岛屿。
// 时间复杂度为O(MN) 空间复杂度O(MN)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	var count int
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i < 0 || j < 0 || i >= r || j >= c || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}

// 广度优先遍历
// 时间复杂度为O(MN) 空间复杂度O(min(M,N))
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	var count int
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = '0'
				queue := make([]int, 0)
				queue = append(queue, i*c+j)
				for len(queue) > 0 {
					num := queue[0]
					queue = queue[1:]
					i, j := num/c, num%c
					if i-1 >= 0 && grid[i-1][j] == '1' {
						queue = append(queue, c*(i-1)+j)
						grid[i-1][j] = '0'
					}
					if i+1 < r && grid[i+1][j] == '1' {
						queue = append(queue, c*(i+1)+j)
						grid[i+1][j] = '0'
					}
					if j-1 >= 0 && grid[i][j-1] == '1' {
						queue = append(queue, c*i+j-1)
						grid[i][j-1] = '0'
					}
					if j+1 < c && grid[i][j+1] == '1' {
						queue = append(queue, c*i+j+1)
						grid[i][j+1] = '0'
					}
				}
			}
		}
	}
	return count
}

// 并查集
