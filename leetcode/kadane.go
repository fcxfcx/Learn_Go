package leetcode

import "math"

// 最大子数和
func MaxSubArray(nums []int) int {
	sum := 0
	ans := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum = nums[i] + max(sum, 0)
		ans = max(ans, sum)
	}
	return ans
}

// 环形子数组的最大和
func MaxSubarraySumCircular(nums []int) int {
	sum1, sum2, sumAll := 0, 0, 0
	ans1, ans2 := math.MinInt32, math.MaxInt32
	// 如果最大和没有利用到环
	for i := 0; i < len(nums); i++ {
		sum1 = nums[i] + max(sum1, 0)
		sumAll += nums[i]
		ans1 = max(ans1, sum1)
	}
	// 如果最大和用到了环，相当于寻找连续的最小子数组，然后用总和减去它
	for j := 0; j < len(nums)-1; j++ {
		sum2 = nums[j] + min(sum2, 0)
		ans2 = min(ans2, sum2)
	}
	return max(ans1, sumAll-ans2)
}
