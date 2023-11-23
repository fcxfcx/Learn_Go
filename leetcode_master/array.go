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
