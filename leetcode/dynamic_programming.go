package leetcode

// 爬楼梯
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// a和b表示前面一级和前前一级的结果
	// 跳到第一级有一种，第二级有两种
	a, b, temp := 1, 2, 0
	for i := 2; i < n; i++ {
		temp = a + b
		a, b = b, temp
	}
	return temp
}

// 打家劫舍
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	a, b, temp := nums[0], nums[1], 0
	for i := 2; i < n; i++ {
		temp = max(a+nums[i], b)
		a, b = b, temp
	}
	return temp
}
