package top_100_liked

// 岛屿数量
func NumIslands(grid [][]byte) int {
	count := 0
	width, height := len(grid[0]), len(grid)
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= width || y < 0 || y >= height || grid[y][x] != '1' {
			return
		}
		grid[y][x] = '0'
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(j, i)
			}
		}
	}
	return count
}

// 腐烂
func OrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{}
	count := 0 // 新鲜橘子数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				// 腐烂橘子入队
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				// 新鲜橘子数量+1
				count++
			}
		}
	}
	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	round := 0 // 分钟数
	for count > 0 && len(queue) != 0 {
		length := len(queue) // 当前轮次的腐烂橘子数量
		for i := 0; i < length; i++ {
			temp := queue[i]
			for _, d := range direction {
				x, y := temp[0]+d[0], temp[1]+d[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 1 {
					continue
				}
				// 新增腐烂橘子
				grid[x][y] = 2
				count--
				queue = append(queue, []int{x, y})
			}
		}
		round++
		queue = queue[length:]
	}
	if count > 0 {
		return -1
	} else {
		return round
	}
}
