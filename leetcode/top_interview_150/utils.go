package top_interview_150

// ----------- 私有工具类方法 --------------
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func max(args ...int) int {
	max := args[0]
	for _, item := range args {
		if item > max {
			max = item
		}
	}
	return max
}

func min(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}

func repeat(num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += " "
	}
	return result
}
