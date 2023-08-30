package leetcode

import "math"

// 搜索插入位置
func SearchInsert(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 搜索二维矩阵
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 首先锁定行
	bottom, top := 0, m
	for bottom < top {
		// 终止条件是bottom == top，区间是左闭右开区间
		mid := (top-bottom)/2 + bottom
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			// mid处大于target了则肯定不在我们需要的区间里
			top = mid
		} else {
			// mid处小于target则还可以继续右移
			bottom = mid + 1
		}
	}
	// 上面判断的是左边界，可以理解为有多少个数小于target
	// 需要锁定的是不大于target的最大索引，因此对于索引来说，结果需要减一
	// 但是需要添加判断避免数组越界
	if bottom-1 < 0 {
		bottom = 0
	} else {
		bottom -= 1
	}
	// 然后锁定列
	left, right := 0, n-1
	for left <= right {
		mid := (right-left)/2 + left
		if matrix[bottom][mid] == target {
			return true
		} else if matrix[bottom][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// 寻找峰值
func FindPeakElement(nums []int) int {
	n := len(nums)
	// 使用开区间
	left, right := -1, n
	get := func(i int) int {
		if i == -1 || i == n+1 {
			return math.MinInt64
		}
		return nums[i]
	}
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		} else if get(mid-1) > get(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

// 搜索旋转排序数组
func SearchRoatedArray(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左边是顺序区间
			if nums[mid] > target && target >= nums[0] {
				// target在0和mid之间，则在左边区间找
				right = mid - 1
			} else {
				// 否则就要去右边区间找
				left = mid + 1
			}
		} else {
			// 右边是顺序区间
			if nums[mid] < target && target <= nums[n-1] {
				// 如果mid处数字小于target，则在右边区间找
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
