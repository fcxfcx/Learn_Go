package leetcode

// 有效的数独
func IsValidSudoku(board [][]byte) bool {
	// 判断横竖
	for i := 0; i < 9; i++ {
		hashmap1 := make(map[byte]bool)
		hashmap2 := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			cur1, cur2 := board[i][j], board[j][i]
			if cur1 != '.' {
				if hashmap1[cur1] {
					return false
				}
				hashmap1[cur1] = true
			}
			if cur2 != '.' {
				if hashmap2[cur2] {
					return false
				}
				hashmap2[cur2] = true
			}
		}
	}
	// 判断九宫格
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			hashmap3 := make(map[byte]bool)
			for _, row := range board[i : i+3] {
				for _, value := range row[j : j+3] {
					if value != '.' && hashmap3[value] {
						return false
					}
					hashmap3[value] = true
				}
			}
		}
	}
	return true
}

// 螺旋矩阵
func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	operations := []string{"right", "down", "left", "up"}
	operation, count := 0, 0
	result := make([]int, 0)
	if m == 1 {
		result = matrix[0]
		return result
	} else if n == 1 {
		for _, value := range matrix {
			result = append(result, value[0])
		}
		return result
	}
	for i, j := 0, 0; count < m*n; {
		result = append(result, matrix[i][j])
		count++
		// 题目条件元素大小在正负100之间，设置为-101以标记为已读取
		matrix[i][j] = -101
		if j == n-1 && i == 0 {
			//右上角
			operation++
			i++
			continue
		} else if i == m-1 && j == n-1 {
			// 右下角
			operation++
			j--
			continue
		} else if i == m-1 && j == 0 {
			// 左下角
			operation++
			i--
			continue
		}
		switch operations[operation%4] {
		case "right":
			if matrix[i][j+1] == -101 {
				operation++
				i++
			} else {
				j++
			}
		case "down":
			if matrix[i+1][j] == -101 {
				operation++
				j--
			} else {
				i++
			}
		case "left":
			if matrix[i][j-1] == -101 {
				operation++
				i--
			} else {
				j--
			}
		case "up":
			if matrix[i-1][j] == -101 {
				operation++
				j++
			} else {
				i--
			}
		}
	}
	return result
}

// 旋转图像
func RotateMatrix(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// 矩阵置零
func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, column := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			column = true
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row = true
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if column {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// 生命游戏
func GameOfLife(board [][]int) {
	// 一共四个状态
	// 1 保持存活
	// 0 保持死亡
	// 2 死变成活
	// -1 活变成死
	m, n := len(board), len(board[0])
	x := [8]int{0, 0, 1, 1, 1, -1, -1, -1}
	y := [8]int{1, -1, 1, 0, -1, -1, 0, 1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := board[i][j]
			count := 0
			for k := 0; k < 8; k++ {
				if i+x[k] < 0 || i+x[k] >= m || j+y[k] < 0 || j+y[k] >= n {
					// 跳过边界情况
					continue
				}
				temp := board[i+x[k]][j+y[k]]
				if temp == 1 || temp == -1 {
					count++
				}
			}
			if cur == 1 {
				if count < 2 {
					board[i][j] = -1
				} else if count > 3 {
					board[i][j] = -1
				}
			} else if cur == 0 && count == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == -1 {
				board[i][j] = 0
			}
		}
	}
}
