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
func MergeInterval(intervals [][]int) [][]int {
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

// 插入区间
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	length := len(intervals)
	if length == 0 {
		result = append(result, newInterval)
		return result
	}
	i := 0
	for i < length && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	for i < length && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)
	for i < length {
		result = append(result, intervals[i])
		i++
	}
	return result
}

// 用最小数量的箭射爆气球
func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	maxRight := points[0][1]
	for _, interval := range points {
		if interval[0] > maxRight {
			count++
			maxRight = interval[1]
			continue
		} 
	}
	return count
}
