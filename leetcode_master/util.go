package leetcode_master

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	maxValue := nums[0]
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
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
