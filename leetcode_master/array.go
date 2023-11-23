package leetcode_master

import "math"

// No.704 二分查找
func Search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// No.27 移除元素
func RemoveElement(nums []int, val int) int {
	n := len(nums)
	left, right := 0, 0
	for right < n {
		if nums[right] == val {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}
	return left + 1
}

// No.977 有序数组的平方
func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right, cur := 0, n-1, n-1
	for left <= right {
		num_left, num_right := nums[left], nums[right]
		if num_left < 0 {
			num_left = -num_left
		}
		if num_right < 0 {
			num_right = -num_right
		}
		if num_left < num_right {
			result[cur] = num_right * num_right
			right--
		} else {
			result[cur] = num_left * num_left
			left++
		}
		cur--
	}
	return result
}

// No.209 长度最小的子数组
func MinSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	n := len(nums)
	minLen, curSum := math.MaxInt64, nums[0]
	for left <= right {
		if curSum < target {
			if right == n-1 {
				break
			} else {
				right++
				curSum += nums[right]
			}
		} else {
			minLen = min(minLen, (right - left + 1))
			curSum -= nums[left]
			left++
		}
	}
	if minLen == math.MaxInt64 {
		return 0
	} else {
		return minLen
	}
}

// No. 59 螺旋矩阵Ⅱ
func GenerateMatrix(n int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		temp := make([]int, n)
		result = append(result, temp)
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex := 0
	row, col := 0, 0
	for i := 1; i <= n*n; i++ {
		result[row][col] = i
		maybeNextRow, maybeNextCol := (row + directions[dirIndex][0]), (col + directions[dirIndex][1])
		if maybeNextRow < 0 || maybeNextRow == n || maybeNextCol < 0 || maybeNextCol == n || result[maybeNextRow][maybeNextCol] != 0 {
			dirIndex = (dirIndex + 1) % 4
		}
		row, col = (row + directions[dirIndex][0]), (col + directions[dirIndex][1])
	}
	return result
}
