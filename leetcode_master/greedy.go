package leetcode_master

import (
	"math"
	"sort"
)

// No.455 分发饼干
func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}

// No.376 摆动序列
func WiggleMaxLength(nums []int) int {
	// 两个节点的差值，记录前一个和当前的
	preDiff, curDiff := 0, 0
	// 默认最后一个数也是峰
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		// 找到一个峰或谷
		if (preDiff >= 0 && curDiff < 0) || (preDiff <= 0 && curDiff > 0) {
			result++
			preDiff = curDiff
		}
	}
	return result
}

// No.53 最大子数组和
func MaxSubArray(nums []int) int {
	result := math.MinInt
	curSum := 0
	for _, num := range nums {
		curSum += num
		if curSum > result {
			result = curSum
		}
		if curSum <= 0 {
			curSum = 0
		}
	}
	return result
}

// No.122 买卖股票的最佳时机Ⅱ
func MaxProfit(prices []int) int {
	total := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			total += (prices[i] - prices[i-1])
		}
	}
	return total
}

// No.55 跳跃游戏
func CanJump(nums []int) bool {
	maxCover := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxCover && i+nums[i] > maxCover {
			maxCover = i + nums[i]
		}
	}
	return maxCover >= len(nums)
}

// No.45 跳跃游戏Ⅱ
func Jump(nums []int) int {
	totalStep := 0
	for tempIndex := 0; tempIndex < len(nums); {
		if tempIndex+nums[tempIndex] >= len(nums)-1 {
			break
		}
		nextIndex := tempIndex
		for i := tempIndex; i < tempIndex+nums[tempIndex]; i++ {
			if i+nums[i] > nextIndex+nums[nextIndex] {
				nextIndex = i
			}
		}
		totalStep++
		tempIndex = nextIndex
	}
	return totalStep
}

// No.1005 K次取反后最大化的数组和
func LargestSumAfterKNegations(nums []int, k int) int {
	total := 0
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
	}
	if k%2 == 1 {
		nums[0] = -nums[0]
	}
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// No.134 加油站
func CanCompleteCircuit(gas []int, cost []int) int {
	totalRest, tempRest := 0, 0
	index := 0
	for i := 0; i < len(gas); i++ {
		totalRest += (gas[i] - cost[i])
		tempRest += (gas[i] - cost[i])
		if tempRest <= 0 {
			index = i + 1
			tempRest = 0
		}
	}
	if totalRest < 0 {
		return -1
	} else {
		return index
	}
}

// No.135 分糖果
func Candy(ratings []int) int {
	n := len(ratings)
	// 总消耗的糖果数
	candies := make([]int, n)
	for index, _ := range candies {
		candies[index] = 1
	}
	// 从左到右保证分更高的有更多糖果
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// 从右到左保证分更高的有更多糖果
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i+1]+1, candies[i])
		}
	}
	total := 0
	for _, num := range candies {
		total += num
	}
	return total
}

// No.860 柠檬水找零
func LemonadeChange(bills []int) bool {
	// 用一个数组存储当前5、10美元的数量
	changes := [2]int{0, 0}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			changes[0] += 1
		} else if bills[i] == 10 {
			changes[1] += 1
			if changes[0] == 0 {
				return false
			} else {
				changes[0] -= 1
			}
		} else {
			if changes[1] > 0 && changes[0] > 0 {
				changes[1] -= 1
				changes[0] -= 1
			} else if changes[0] >= 3 {
				changes[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
