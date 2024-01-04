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
