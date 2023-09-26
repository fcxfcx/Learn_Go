package top_interview_150

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

// 在排序数组中查找元素的第一个和最后一个位置
func SearchRange(nums []int, target int) []int {
	// 搜索左边界
	left := binarySerch(nums, target)
	// 搜索比target大1的左边界，下标减一即为target的右边界
	right := binarySerch(nums, target+1)
	if left == len(nums) || nums[left] != target {
		// 如果不含有target
		return []int{-1, -1}
	} else {
		return []int{left, right - 1}
	}
}

func binarySerch(nums []int, target int) int {
	// 使用二分查找在升序数组寻找第一个大于等于target的下标（左边界）
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			// 左移右边界
			right = mid - 1
		} else {
			// 右移左边界
			left = mid + 1
		}
	}
	return left
}

// 寻找旋转数组中的最小值
func FindMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[n-1] {
			// 如果中点小于最右侧，那么右边区域可以排除
			right = mid
		} else {
			// 否则的话说明左边区域可以排除
			left = mid + 1
		}
	}
	return nums[left]
}

// 寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	half := (m + n) / 2
	if (m+n)%2 == 0 {
		// 如果两个数组长度和为偶数，则需要寻找的是half和half-1处的两个数字(对应数组索引)
		// 索引映射到第k小的数，需要加一
		num1 := getKthMinNumber(nums1, nums2, half+1)
		num2 := getKthMinNumber(nums1, nums2, half)
		// 注意这里不能先除再转float否则会在除法部分就丢弃小数部分
		return float64(num1+num2) / 2.0
	} else {
		// 如果两个数组长度和为奇数，则需要寻找的是half处的数字(对应数组索引)
		// 索引映射到第k小的数，需要加一
		return float64(getKthMinNumber(nums1, nums2, half+1))
	}
}

func getKthMinNumber(nums1, nums2 []int, k int) int {
	// 获取两个数组中第k小的数字
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			// 如果第一个数组已经排除完毕，则直接从第二个数组取
			// 取的是从当前索引开始的第k个数字
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			// 对于第二个数组同理
			return nums1[index1+k-1]
		}
		if k == 1 {
			// 如果取最小的数，则比较两个数组的第一个数
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		// 确定两个数组往后第k/2个数字，同时避免越界
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		if nums1[newIndex1] >= nums2[newIndex2] {
			// 如果第一个数组的数字大，则第二个数组的前部分可以排除
			// 排除后，需要寻找的数字数量也减少对应的数目
			k -= (newIndex2 - index2) + 1
			index2 = newIndex2 + 1
		} else {
			// 反之亦然
			k -= (newIndex1 - index1) + 1
			index1 = newIndex1 + 1
		}
	}
}
