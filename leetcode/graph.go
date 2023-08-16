package leetcode

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

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

// 克隆图
func CloneGraph(node *GraphNode) *GraphNode {
	hashmap := make(map[*GraphNode]*GraphNode)
	var clone func(node *GraphNode) *GraphNode
	clone = func(node *GraphNode) *GraphNode {
		if node == nil {
			return nil
		}
		if item, ok := hashmap[node]; ok {
			return item
		}
		cloneNode := &GraphNode{
			Val:       node.Val,
			Neighbors: []*GraphNode{},
		}
		hashmap[node] = cloneNode
		for _, neighbor := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, clone(neighbor))
		}
		return cloneNode
	}
	return clone(node)
}

// 除法求值
func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 第一步：将字符串转为数字，方便构建图
	hash := make(map[string]int, 0)
	id := 0
	for _, e := range equations {
		// 条件中的两个数，a为被除数，b为除数
		// 如果没有构建数字的id则新建
		a, b := e[0], e[1]
		if _, ok := hash[a]; !ok {
			hash[a] = id
			id++
		}
		if _, ok := hash[b]; !ok {
			hash[b] = id
			id++
		}
	}

	// 第二步：构建图
	// 此时id的大小代表一共有多少个字符串出现在条件里，可以作为图的边长
	graph := make([][]float64, id)
	for i := 0; i < id; i++ {
		graph[i] = make([]float64, id)
		// 自己除以自己得1
		graph[i][i] = 1
	}
	for index, value := range values {
		// 条件中的两个数，a为被除数，b为除数
		a, b := equations[index][0], equations[index][1]
		// 从哈希表里提取它们的数字id，作为图的坐标
		graph[hash[a]][hash[b]] = value
		graph[hash[b]][hash[a]] = 1 / value
	}

	// 第三步：DFS
	// 用一个哈希表记录某个数是否已经被访问过了
	visited := map[int]bool{}
	temp := 0.
	var dfs func(index int, target int, value float64)
	dfs = func(index, target int, value float64) {
		if index == target {
			temp = value
			return
		}
		for i := 0; i < id; i++ {
			// 在index启示作为被除数的地方开始，深度优先搜索尝试走通至除数
			// 此处设置visited目的是避免无限循环
			if graph[index][i] != 0 && !visited[i] {
				visited[i] = true
				dfs(i, target, value*graph[index][i])
			}
		}
	}

	// 第四步：对需要查询的结果进行搜索
	res := []float64{}
	for _, query := range queries {
		temp = -1.
		// 首先判断是否有被除数和除数对应的id，如果没有的话说明算不出结果，直接输出-1
		if _, ok := hash[query[0]]; ok {
			if _, ok1 := hash[query[1]]; ok1 {
				// 如果都有对应的id说明可以尝试进行运算，如果中途无法走通则同样返回-1
				// 开始之前需要将visited重置
				visited = map[int]bool{}
				dfs(hash[query[0]], hash[query[1]], 1)
			}
		}
		res = append(res, temp)
	}
	return res
}

// 课程表
func CanFinish(numCourses int, prerequisites [][]int) bool {
	valid := true
	// 用图来表示学某一课程的前置条件
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		// 如果要学a课程必须先学b课程
		graph[p[0]] = append(graph[p[0]], p[1])
	}
	// searched为0代表未搜寻，为1代表正在搜寻，为2代表已搜寻
	searched := make([]int, numCourses)
	var dfs func(index int)
	dfs = func(index int) {
		// 将当前搜索的课程置为正在搜寻
		searched[index] = 1
		for _, c := range graph[index] {
			if searched[c] == 1 {
				// 如果探索到了正在搜寻的课程，则说明成环了，无法上完所有课程
				valid = false
				return
			}
			if searched[c] == 0 {
				// 对于未搜寻的课程则开始搜寻
				dfs(c)
				if !valid {
					return
				}
			}
		}
		searched[index] = 2
	}
	// 保证整个图不成环
	for i := 0; i < numCourses && valid; i++ {
		if searched[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// 课程表Ⅱ
func FindOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	// 构建邻接表
	afterMp := map[int][]int{}
	// 储存每个课程需要多少前置课
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		// 后修课程和先修课程
		after, pre := p[0], p[1]
		afterMp[pre] = append(afterMp[pre], after)
		indegree[after] += 1
	}
	// 使用队列来从前置课开始搜寻
	queue := []int{}
	// 学过的课程数量
	alreadyLearned := 0
	// 将无前置课的课程先初始化到队列中
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	// 不断读取队列中无需前置课的课程
	for len(queue) > 0 {
		alreadyLearned++
		temp := queue[0]
		result = append(result, temp)
		queue = queue[1:]
		for _, after := range afterMp[temp] {
			// 学习了前置课后，对应的后修课程所需前置课数量减一
			indegree[after]--
			if indegree[after] == 0 {
				// 如果后修课程已无前置课程，则入队
				queue = append(queue, after)
			}
		}
	}
	if alreadyLearned == numCourses {
		return result
	} else {
		return nil
	}
}

func SnakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	// 构造一个结构体存储（当前id，当前使用步数）
	type pair struct{ index, step int }
	// 构造队列，加入起点
	queue := []pair{{1, 0}}
	for len(queue) > 0 {
		// 出队操作
		tempPair := queue[0]
		queue = queue[1:]
		// 遍历所有走法
		for i := 1; i <= 6; i++ {
			tempID := tempPair.index + i
			if tempID > n*n {
				// 超出边界的情况
				break
			}
			r, c := idToRC(tempID, n)
			if board[r][c] > 0 {
				// 存在蛇或者梯子的情况
				tempID = board[r][c]
			}
			if tempID == n*n {
				// 到达终点
				return tempPair.step + 1
			}
			if !visited[tempID] {
				visited[tempID] = true
				queue = append(queue, pair{tempID, tempPair.step + 1})
			}
		}
	}
	// 如果到达不了返回-1
	return -1
}

func idToRC(index int, n int) (r, c int) {
	// 将蛇形棋盘中的ID转换为矩阵中的行列
	r = (index - 1) / n
	c = (index - 1) % n
	if r%2 == 1 {
		// 如果是偶数排需要倒序
		c = n - c - 1
	}
	// 棋盘本身是从下到上排列的
	r = n - r - 1
	return
}
