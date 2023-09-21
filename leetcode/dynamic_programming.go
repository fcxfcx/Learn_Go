package leetcode

import "math"

// 爬楼梯
func ClimbStairs(n int) int {
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
func Rob(nums []int) int {
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

// 单词拆分
func WordBreak(s string, wordDict []string) bool {
	dic := make(map[string]bool, 0)
	maxWord := 0
	for _, word := range wordDict {
		if len(word) > maxWord {
			// 顺便维护一下词典里最长的词是多少
			maxWord = len(word)
		}
		dic[word] = true
	}
	// dp[i]为true代表i之前的都可以组成
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if i-j > maxWord {
				continue
			}
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 零钱兑换
func CoinChange(coins []int, amount int) int {
	//dp[i]代表面值i以前的钱最少兑换的硬币数
	dp := make([]int, amount+1)
	for i := range dp {
		// 方便之后判断最小值，除以二是为了避免+1后溢出
		dp[i] = math.MaxInt / 2
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			// 枚举最后一枚硬币的面值
			target := i - coins[j]
			if target < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[target]+1)
		}
	}
	if dp[amount] < math.MaxInt/2 {
		return dp[amount]
	}
	return -1
}

// 最长递增子序列
func LengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	res := 0
	tail := make([]int, len(nums))
	tail[0] = nums[0]
	for _, num := range nums {
		i, j := 0, res
		for i < j {
			// 使用二分法查找当前数应该插入的位置
			m := (i + j) / 2
			if num > tail[m] {
				i = m + 1
			} else {
				j = m
			}
		}
		tail[i] = num
		if j == res {
			// 如果当前数插入的地方是最后，则代表结果加一
			res += 1
		}
	}
	return res
}

// 三角形的最短路径和
func MinimumTotal(triangle [][]int) int {
	layer := len(triangle)
	if layer == 1 {
		return triangle[0][0]
	}
	dp := make([]int, layer+1)
	for i := layer - 1; i >= 0; i-- {
		// 从下往上推导
		tempT := triangle[i]
		for j := 0; j < len(tempT); j++ {
			// 往上推导一层，从两个可连接路径中选取更小的那个
			dp[j] = min(dp[j], dp[j+1]) + tempT[j]
		}
	}
	return dp[0]
}

// 最大路径和
func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 使用原数组，减少空间消耗
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				// 上边界，只能从左来
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				// 左边界，只能从上来
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

// 不同路径Ⅱ
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				// 障碍格直接跳过
				continue
			} else if i == 0 && j == 0 {
				// 初始格子有一条路径，用负数和障碍进行区分
				obstacleGrid[i][j] = -1
			} else if i == 0 {
				// 上边界，只能从左边过来
				if obstacleGrid[i][j-1] != 1 {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				}
			} else if j == 0 {
				// 左边界，只能从上来
				if obstacleGrid[i-1][j] != 1 {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				}
			} else {
				// 当前格子不是障碍物或者边界
				if obstacleGrid[i-1][j] != 1 {
					obstacleGrid[i][j] += obstacleGrid[i-1][j]
				}
				if obstacleGrid[i][j-1] != 1 {
					obstacleGrid[i][j] += obstacleGrid[i][j-1]
				}
			}
		}
	}
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	// 前面都是用负数进行累加的，所以要返回负数结果
	return -obstacleGrid[m-1][n-1]
}

// 最长回文子串
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	expand := func(left, right int) (int, int) {
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		}
		// 边界条件是最后一个不属于回文子串的索引，所以要去掉
		return left + 1, right - 1
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 这里右边界可以超，因为expand方法里有边界判断
		left1, right1 := expand(i, i)
		left2, right2 := expand(i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// 交错字符串
func IsInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// 用滚动数组节省空间
	dp := make([]bool, n2+1)
	// 最初始的字符串可以认为是空字符串，所以一定符合
	dp[0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			// 对应的s3的下标
			temp := i + j - 1
			if i > 0 {
				// 如果由dp[i-1][j]变为dp[i][j]，也就是在s1中多选一个字符
				// 则符合条件的要求是当前的dp[j]符合，并且新的字符和s3的下一个字符相同
				// 注意dp的下标对应的是字符串真实下标加一的结果
				dp[j] = dp[j] && s1[i-1] == s3[temp]
			}
			if j > 0 {
				// 如果由dp[i][j-1]变为dp[i][j]，也就是在s2中多选一个字符
				// 如果上一步已经符合，则不用再判断，用上一步的结果即可
				// 否则符合的条件是之前的dp[j-1]符合，并且新字符串与s3的下一个字符相同
				dp[j] = dp[j] || (dp[j-1] && s2[j-1] == s3[temp])
			}
		}
	}
	// 注意下标是加过1的
	return dp[n2]
}
