package leetcode_master

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
