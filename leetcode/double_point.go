package leetcode

import (
	"regexp"
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
