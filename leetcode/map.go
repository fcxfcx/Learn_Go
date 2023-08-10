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

// 被围绕的区域
func Solve(board [][]byte) {
	height, width := len(board), len(board[0])
	var dfs func(board [][]byte, i int, j int)
	dfs = func(board [][]byte, i, j int) {
		if i < 0 || i >= height || j < 0 || j >= width || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#'
		dfs(board, i+1, j)
		dfs(board, i-1, j)
		dfs(board, i, j+1)
		dfs(board, i, j-1)
	}
	for i := 0; i < height; i++ {
		dfs(board, i, 0)
		dfs(board, i, width-1)
	}
	for j := 0; j < width; j++ {
		dfs(board, 0, j)
		dfs(board, height-1, j)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
