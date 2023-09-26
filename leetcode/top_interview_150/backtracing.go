package top_interview_150

import "sort"

// 电话号码的字母组合
func LetterCombinations(digits string) []string {
	words := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	if len(digits) == 0 {
		return []string{}
	}
	var dfs func(digits string) []string
	dfs = func(digits string) []string {
		result := make([]string, 0)
		if len(digits) == 1 {
			// 最后一位
			for _, word := range words[digits[0]-'2'] {
				result = append(result, string(word))
			}
		} else {
			for _, word := range words[digits[0]-'2'] {
				tempWord := string(word)
				for _, afterString := range dfs(digits[1:]) {
					result = append(result, tempWord+afterString)
				}
			}
		}
		return result
	}
	return dfs(digits)
}

// 组合
func Combine(n int, k int) (ans [][]int) {
	path := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) > k || len(path)+(n-i+1) < k {
			// 如果已经装填完了或者后面的数字不够装填
			return
		}
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for next := i; next <= n; next++ {
			path = append(path, next)
			dfs(next)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return
}

// 全排列
func Permute(nums []int) (ans [][]int) {
	hashset := make(map[int]bool, 0)
	for _, num := range nums {
		hashset[num] = true
	}
	path := []int{}
	var dfs func(length int)
	dfs = func(length int) {
		if length == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for num := range hashset {
			if hashset[num] {
				path = append(path, num)
				hashset[num] = false
				dfs(length - 1)
				path = path[:len(path)-1]
				hashset[num] = true
			}
		}
	}
	dfs(len(nums))
	return ans
}

// 组合总合
func CombinationSum(candidates []int, target int) (ans [][]int) {
	path := []int{}
	// 升序排序
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var dfs func(start int)
	dfs = func(start int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
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

// N 皇后
func SolveNQueens(n int) (ans [][]string) {
	// columns代表哪些列已经站了皇后
	columns := make(map[int]bool, 0)
	// diagonals1和2代表两个对角线是否站了皇后
	// 左上到右下，行列之差相等，左下到右上，行列之和相等，因此可以用一个数表示一条线
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	// 每一行肯定都会放皇后，所以可以用一个数组表示每行放在哪一列上面
	queens := make([]int, n)

	var backtrace func(row int)
	backtrace = func(row int) {
		if row == n {
			// 放满了则找到了一个可行的解
			board := generateBoard(queens, n)
			ans = append(ans, board)
			return
		}
		// 遍历当前行，寻找可能的解
		for i := 0; i < n; i++ {
			if columns[i] {
				// 同列有放皇后
				continue
			}
			if diagonals1[row-i] {
				// 左上到右下对角线有放皇后
				continue
			}
			if diagonals2[row+i] {
				// 左下到右上对角线有放皇后
				continue
			}
			// 都没有放则可以放当前位置
			queens[row] = i
			columns[i] = true
			diagonals1[row-i] = true
			diagonals2[row+i] = true
			// 继续考虑下一行
			backtrace(row + 1)
			// 回溯后将当前位置放置的皇后取消
			queens[row] = -1
			delete(columns, i)
			delete(diagonals1, row-i)
			delete(diagonals2, row+i)
		}
	}
	backtrace(0)
	return
}

func generateBoard(queens []int, n int) (board []string) {
	for _, queen := range queens {
		row := make([]byte, n)
		for i := 0; i < n; i++ {
			row[i] = '.'
		}
		row[queen] = 'Q'
		board = append(board, string(row))
	}
	return
}

// N皇后Ⅱ
func TotalNQueens(n int) int {
	// 此题是N皇后的简单版本，只需要判读有几个解就行了，因此稍微改动代码
	total := 0
	// columns代表哪些列已经站了皇后
	columns := make(map[int]bool, 0)
	// diagonals1和2代表两个对角线是否站了皇后
	// 左上到右下，行列之差相等，左下到右上，行列之和相等，因此可以用一个数表示一条线
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}

	var backtrace func(row int)
	backtrace = func(row int) {
		if row == n {
			total += 1
			return
		}
		// 遍历当前行，寻找可能的解
		for i := 0; i < n; i++ {
			if columns[i] {
				// 同列有放皇后
				continue
			}
			if diagonals1[row-i] {
				// 左上到右下对角线有放皇后
				continue
			}
			if diagonals2[row+i] {
				// 左下到右上对角线有放皇后
				continue
			}
			// 都没有放则可以放当前位置
			columns[i] = true
			diagonals1[row-i] = true
			diagonals2[row+i] = true
			// 继续考虑下一行
			backtrace(row + 1)
			// 回溯后将当前位置放置的皇后取消
			delete(columns, i)
			delete(diagonals1, row-i)
			delete(diagonals2, row+i)
		}
	}
	backtrace(0)
	return total
}

// 括号生成
func GenerateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs func(index int, left int)
	dfs = func(index int, left int) {
		if index == m {
			tmp := make([]byte, m)
			copy(tmp, path)
			ans = append(ans, string(tmp))
			return
		}
		if left < n {
			// 可以填左括号
			path[index] = '('
			dfs(index+1, left+1)
		}
		if index-left < left {
			// 可以填右括号
			path[index] = ')'
			dfs(index+1, left)
		}
	}
	dfs(0, 0)
	return
}

// 单词搜索
func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	next := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		ch := board[i][j]
		if ch != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		board[i][j] = '#'
		for _, new := range next {
			nx := i + new[0]
			ny := j + new[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || board[nx][ny] == '#' {
				// 越界或已选用
				continue
			}
			if dfs(nx, ny, index+1) {
				return true
			}
		}
		board[i][j] = ch
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
