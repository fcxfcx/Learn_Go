package leetcode_master

import "math"

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

// No.1049 最后一块石头的重量II
func LastStoneWeightII(stones []int) int {
	sum := 0
	dp := [15001]int{}
	for _, num := range stones {
		sum += num
	}
	target := sum / 2
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return (sum - dp[target]) - dp[target]
}

// No.494 目标和
func FindTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果不能被2整除说明不可能有解
	if (sum+target)%2 != 0 {
		return 0
	}
	if abs(target) > sum {
		return 0
	}
	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bag]
}

// No.474
func FindMaxForm(strs []string, m int, n int) int {
	// 把字符串的1和0统计出来
	bitStrs := [][2]int{}
	for _, str := range strs {
		zero, one := 0, 0
		for _, b := range []byte(str) {
			if b-'0' == 0 {
				zero++
			} else {
				one++
			}
		}
		bitStrs = append(bitStrs, [2]int{zero, one})
	}
	// 构造二维背包
	dp := make([][]int, m+1)
	for index := range dp {
		dp[index] = make([]int, n+1)
	}
	// 01背包问题
	for _, bitStr := range bitStrs {
		for i := m; i >= bitStr[0]; i-- {
			for j := n; j >= bitStr[1]; j-- {
				dp[i][j] = max(dp[i][j], dp[i-bitStr[0]][j-bitStr[1]]+1)
			}
		}
	}
	return dp[m][n]
}

// No.518 零钱兑换Ⅱ
func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// No.279 完全平方数
func NumSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

// No.139 单词拆分
func WordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool, 0)
	for _, word := range wordDict {
		hash[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && hash[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// No.198 打家劫舍
func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	// 使用dp数组储存每一家的最大收益
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

// No.213 打家劫舍Ⅱ
func RobTwo(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	result1 := Rob(nums[1:])
	result2 := Rob(nums[0 : n-1])
	return max(result1, result2)
}

// No.337 打家劫舍Ⅲ
func RobTree(root *TreeNode) int {
	var traversal func(node *TreeNode) [2]int
	traversal = func(node *TreeNode) [2]int {
		dp := [2]int{}
		if node == nil {
			return [2]int{0, 0}
		}
		left := traversal(node.Left)
		right := traversal(node.Right)
		// 劫当前节点
		dp[0] = node.Val + left[1] + right[1]
		// 不劫当前节点
		dp[1] = max(left[0], left[1]) + max(right[0], right[1])
		return dp
	}
	result := traversal(root)
	return max(result[0], result[1])
}

// No.121 买卖股票的最佳时机
func ProfitOne(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

// No.122 买卖股票的最佳时机Ⅱ
func ProfitTwo(prices []int) int {
	total := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			total += prices[i] - prices[i-1]
		}
	}
	return total
}

// No.123 买卖股票的最佳时机Ⅲ
func ProfitThree(prices []int) int {
	dp := make([][5]int, len(prices))
	// 每天有五种状态：
	// 1. 什么都不做
	// 2. 第一次购入
	// 3. 第一次售出
	// 4. 第二次购入
	// 5. 第二次售出
	dp[0] = [5]int{0, -prices[0], 0, -prices[0], 0}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		// 当天第一次购入状态有两种可能，即当天购入，和延用前一天已经购入的情况
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		// 当天第一次售出状态有两种可能，即当天售出，和延用前一天售出的状态
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = max(dp[i-1][2]-prices[i], dp[i-1][3])
		dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])
	}
	return dp[len(prices)-1][4]
}
