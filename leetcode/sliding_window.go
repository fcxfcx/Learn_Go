package leetcode

import (
	"math"
)

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

// 串联所有单词的子串
func FindSubstring(s string, words []string) []int {
	n, m := len(words), len(words[0])
	result := make([]int, 0)
	slen := len(s)
	hashmap1 := make(map[string]int)
	for _, value := range words {
		hashmap1[value] += 1
	}
	for i := 0; i < m; i++ {
		var count = 0
		hashmap2 := make(map[string]int)
		for l, r := i, i; r <= slen-m; r += m {
			word := s[r : r+m]
			if num, found := hashmap1[word]; found {
				for hashmap2[word] >= num {
					hashmap2[s[l:l+m]]--
					l += m
					count--
				}
				hashmap2[word]++
				count++
			} else {
				for l < r {
					hashmap2[s[l:l+m]]--
					l += m
					count--
				}
				// 外圈的r会加m，所以这里l等于r之后还要加上一个m
				l += m
			}
			if count == n {
				result = append(result, l)
			}
		}
	}
	return result
}

// 最小覆盖子串
func MinWindow(s string, t string) string {
	slen, tlen := len(s), len(t)
	result := ""
	// 特判
	if slen < tlen {
		return result
	}
	hashmap1 := make(map[byte]int)
	for _, b := range t {
		hashmap1[byte(b)]++
	}

	count := 0
	hashmap2 := make(map[byte]int)
	for l, r := 0, 0; r < slen && l <= r; r++ {
		char := byte(s[r])
		if num, found := hashmap1[char]; found {
			hashmap2[char] += 1
			if hashmap2[char] <= num {
				count++
			}
		}
		if count == tlen {
			for l < r {
				// 移动左侧直至不满足t
				num, ok := hashmap2[s[l]]
				if ok {
					if num == hashmap1[s[l]] {
						break
					} else {
						hashmap2[s[l]] -= 1
						l += 1
					}
				} else {
					l += 1
				}
			}
			if len(result) == 0 || r-l < len(result) {
				result = s[l : r+1]
			}
			hashmap2[s[l]] -= 1
			l += 1
			count--
		}
	}
	return result
}
