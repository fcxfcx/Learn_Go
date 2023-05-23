package leetcode

import (
	"fmt"
	"math"
)

// 合并两个有序数组
func Merge(nums1 []int, m int, nums2 []int, n int) {
	m_index, n_index, cur := m-1, n-1, 0
	for i := m + n - 1; i >= 0; i-- {
		if m_index == -1 {
			cur = nums2[n_index]
			n_index--
		} else if n_index == -1 {
			cur = nums1[m_index]
			m_index--
		} else if nums1[m_index] < nums2[n_index] {
			cur = nums2[n_index]
			n_index--
		} else {
			cur = nums1[m_index]
			m_index--
		}
		nums1[i] = cur
	}
	fmt.Println(nums1)
}

// 移除数组中某一元素
func RemoveElement(nums []int, val int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[head] == val {
			nums[head] = nums[tail]
			tail--
		} else {
			head++
		}
	}
	fmt.Println(nums)
	return head
}

// 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	head, tail := 0, 1
	for tail < len(nums) {
		if nums[head] == nums[tail] {
			tail++
		} else {
			nums[head+1] = nums[tail]
			head++
		}
	}
	fmt.Println(nums)
	return head + 1
}

// 删除有序数组中的重复项Ⅱ 重复项最多可保留两个
func RemoveDuplicatesPlus(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow := 2
	for fast := 2; fast < n; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 返回数组中数量大于一半的元素
func MajorityElement(nums []int) int {
	count, cur := 0, 0
	for _, num := range nums {
		if count == 0 {
			cur = num
			count++
		} else if cur != num {
			count--
		} else {
			count++
		}
	}
	return cur
}

// 轮转数组元素
func Rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// 买卖股票的最佳时机
func MaxProfit(prices []int) int {
	min, profit := math.MaxInt64, 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		profit = int(math.Max(float64(profit), float64(price-min)))
	}
	return profit
}

// 买卖股票的最佳时机Ⅱ（多次买卖）
func MaxProfitPlus(prices []int) int {
	length := len(prices)
	dp := make([][2]int, length)
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}

// 跳跃游戏
func CanJump(nums []int) bool {
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if i > maxLength {
			return false
		} else {
			maxLength = max(maxLength, i+nums[i])
		}
	}
	return true
}

// H指数
func HIndex(citations []int) int {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	for i, total := n, 0; i > 0; i-- {
		total += counter[i]
		if total >= i {
			return i
		}
	}
	return 0
}

// ----------- 私有工具类方法 --------------
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
