package leetcode

import "math"

// 长度最小的子数组
func MinSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	ans := math.MaxInt
	sum, n := 0, len(nums)
	for end < n {
		sum += nums[end]
		if ans == 1 {
			return ans
		}
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

// 无重复字符的最长字串
func LengthOfLongestSubstring(s string) int {
	start, end, n, ans := 0, 0, len(s), 1
	if n == 0 {
		return 0
	}
	hashmap := make(map[byte]bool)
	hashmap[s[start]] = true
	for end < n {
		for end < n-1 && !hashmap[s[end+1]] {
			end++
			hashmap[s[end]] = true
			ans = max(ans, end-start+1)
		}
		if start < end {
			hashmap[s[start]] = false
			start++
		} else {
			start++
			end++
		}
	}
	return ans
}
