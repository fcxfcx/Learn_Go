package leetcode_master

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
