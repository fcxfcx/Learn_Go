package leetcode

import (
	"sort"
	"strconv"
)

// 汇总区间
func SummaryRanges(nums []int) []string {
	length := len(nums)
	result := make([]string, 0)
	if length == 0 {
		return result
	} else if length == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	start, end := 0, 0
	for end < length {
		if end != length-1 && nums[end+1] == nums[end]+1 {
			end++
		} else {
			if start == end {
				result = append(result, strconv.Itoa(nums[start]))
			} else {
				temp := ""
				temp += strconv.Itoa(nums[start])
				temp += "->"
				temp += strconv.Itoa(nums[end])
				result = append(result, temp)
			}
			end++
			start = end
		}
	}
	return result
}

// 合并区间
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := range intervals {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
			continue
		} else if prev[1] < cur[1] {
			prev[1] = cur[1]
		}
	}
	res = append(res, prev)
	return res
}
