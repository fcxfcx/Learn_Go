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
