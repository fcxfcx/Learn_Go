package leetcode

// ----------- 私有工具类方法 --------------
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func repeat(num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += " "
	}
	return result
}
