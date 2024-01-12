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
