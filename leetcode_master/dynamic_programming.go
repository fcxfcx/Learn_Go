package leetcode_master

// No.509 斐波那契数
func Fib(n int) int {
	pre, cur := 0, 1
	next := 0
	if n == 0 {
		return pre
	} else if n == 1 {
		return cur
	}
	for i := 1; i < n; i++ {
		next = pre + cur
		pre = cur
		cur = next
	}
	return cur
}

// No.70 爬楼梯
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// 储存前面两级台阶的走法有多少
	dp := [2]int{}
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[1], dp[0] = dp[1]+dp[0], dp[1]
	}
	return dp[1]
}

// No.746 使用最小花费爬楼梯
func MinCostClimbingStairs(cost []int) int {
	// 储存前面两级台阶最小花费
	dp := [2]int{cost[0], cost[1]}
	for i := 2; i < len(cost); i++ {
		dp[0], dp[1] = dp[1], min(dp[0], dp[1])+cost[i]
	}
	return min(dp[1], dp[0])
}

// No.62 不同路径
func UniquePaths(m int, n int) int {
	// 储存当前行每个格子的走法
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		// 每一行的第一个格子一定只有一种走法
		dp[0] = 1
		for j := 1; j < n; j++ {
			if i == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}
	return dp[n-1]
}
