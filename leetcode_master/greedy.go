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
