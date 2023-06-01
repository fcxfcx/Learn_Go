package leetcode

import (
	"regexp"
	"sort"
	"strings"
)

// 验证回文串
func IsPalindrome(s string) bool {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

// 判断子序列
func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)-1
}

// 两数之和Ⅱ 输入有序数组
func TwoSum(numbers []int, target int) []int {
	index1, index2 := 0, len(numbers)-1
	for index1 < index2 {
		if numbers[index1]+numbers[index2] < target {
			index1++
		} else if numbers[index1]+numbers[index2] > target {
			index2--
		} else {
			break
		}
	}
	return []int{index1 + 1, index2 + 1}
}

// 盛水最多的容器
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxHeight := 0
	for left < right {
		width := right - left
		if height[left] < height[right] {
			maxHeight = max(maxHeight, height[left]*width)
			left++
		} else {
			maxHeight = max(maxHeight, height[right]*width)
			right--
		}
	}
	return maxHeight
}

// 三数之和
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
