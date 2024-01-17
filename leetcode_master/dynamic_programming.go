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

// No.63 不同路径Ⅱ
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
			} else if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}

// No.343 整数拆分
func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, i*(i-j)))
		}
	}
	return dp[n]
}

// No.96 不同的二叉搜索树
func NumTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[i-j] * dp[j-1]
		}
	}
	return dp[n]
}

// No.416 分割等和子集
func CanPartition(nums []int) bool {
	dp := [10001]int{}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[target] == target {
		return true
	} else {
		return false
	}
}
