package leetcode

// 岛屿数量
func NumIslands(grid [][]byte) int {
	count := 0
	height := len(grid)
	width := len(grid[0])
	var dfs func(grid [][]byte, i int, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= height || j < 0 || j >= width || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}
