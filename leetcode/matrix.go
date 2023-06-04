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
