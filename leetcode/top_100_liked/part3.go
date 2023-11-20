package top_100_liked

import "sort"

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

// 课程表
func CanFinish(numCourses int, prerequisites [][]int) bool {
	afterMap := map[int][]int{}         // 储存某一课程的后继课程
	indgrees := make([]int, numCourses) // 记录每个课程所需的前置课程数量
	learned := 0
	for _, v := range prerequisites {
		// 课程v[0] 需要先修课程 v[1]
		afterMap[v[1]] = append(afterMap[v[1]], v[0])
		indgrees[v[0]]++
	}
	q := []int{} // 当前可以修的课程
	for i := 0; i < numCourses; i++ {
		// 无需修前置课程的课程先入队
		if indgrees[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		learned++
		for _, v := range afterMap[temp] {
			// temp对应的后继课程所需前置课程数量减1
			indgrees[v]--
			if indgrees[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return learned == numCourses
}

// 全排列
func Permute(nums []int) (ans [][]int) {
	if len(nums) == 0 {
		return ans
	}
	hashset := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		hashset[nums[i]] = true
	}

	path := []int{}
	var dfs func(length int)
	dfs = func(length int) {
		if length == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for num := range hashset {
			if !hashset[num] {
				continue
			}
			path = append(path, num)
			hashset[num] = false
			dfs(length - 1)
			hashset[num] = true
			path = path[:len(path)-1]
		}
	}
	dfs(len(nums))
	return ans
}

// 子集
func Subsets(nums []int) (ans [][]int) {
	set := []int{}

	// 对于指定长度的子集，某一处的数字只有两种情况即是选取或不选取
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		// 选取当前数字
		set = append(set, nums[cur])
		dfs(cur + 1)
		// 不选取当前数字
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

// 电话号码
func LetterCombinations(digits string) (ans []string) {
	words := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	length := len(digits)
	if length == 0 {
		return
	}
	path := ""
	var dfs func(index int)
	dfs = func(index int) {
		if index == length {
			ans = append(ans, path)
			return
		}
		for _, v := range words[digits[index]-'2'] {
			path += string(v)
			dfs(index + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}

// 组合总和
func CombinationSum(candidates []int, target int) (ans [][]int) {
	path := []int{}
	// 升序排序数组
	sort.Ints(candidates)
	var dfs func(start int)
	dfs = func(start int) {
		if target < 0 {
			return
		}
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			target -= candidates[i]
			dfs(i)
			target += candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
