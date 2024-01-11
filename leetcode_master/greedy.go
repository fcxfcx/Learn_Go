package leetcode_master

import (
	"math"
	"sort"
	"strconv"
)

// No.455 分发饼干
func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}

// No.376 摆动序列
func WiggleMaxLength(nums []int) int {
	// 两个节点的差值，记录前一个和当前的
	preDiff, curDiff := 0, 0
	// 默认最后一个数也是峰
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		// 找到一个峰或谷
		if (preDiff >= 0 && curDiff < 0) || (preDiff <= 0 && curDiff > 0) {
			result++
			preDiff = curDiff
		}
	}
	return result
}

// No.53 最大子数组和
func MaxSubArray(nums []int) int {
	result := math.MinInt
	curSum := 0
	for _, num := range nums {
		curSum += num
		if curSum > result {
			result = curSum
		}
		if curSum <= 0 {
			curSum = 0
		}
	}
	return result
}

// No.122 买卖股票的最佳时机Ⅱ
func MaxProfit(prices []int) int {
	total := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			total += (prices[i] - prices[i-1])
		}
	}
	return total
}

// No.55 跳跃游戏
func CanJump(nums []int) bool {
	maxCover := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxCover && i+nums[i] > maxCover {
			maxCover = i + nums[i]
		}
	}
	return maxCover >= len(nums)
}

// No.45 跳跃游戏Ⅱ
func Jump(nums []int) int {
	totalStep := 0
	for tempIndex := 0; tempIndex < len(nums); {
		if tempIndex+nums[tempIndex] >= len(nums)-1 {
			break
		}
		nextIndex := tempIndex
		for i := tempIndex; i < tempIndex+nums[tempIndex]; i++ {
			if i+nums[i] > nextIndex+nums[nextIndex] {
				nextIndex = i
			}
		}
		totalStep++
		tempIndex = nextIndex
	}
	return totalStep
}

// No.1005 K次取反后最大化的数组和
func LargestSumAfterKNegations(nums []int, k int) int {
	total := 0
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
	}
	if k%2 == 1 {
		nums[0] = -nums[0]
	}
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// No.134 加油站
func CanCompleteCircuit(gas []int, cost []int) int {
	totalRest, tempRest := 0, 0
	index := 0
	for i := 0; i < len(gas); i++ {
		totalRest += (gas[i] - cost[i])
		tempRest += (gas[i] - cost[i])
		if tempRest <= 0 {
			index = i + 1
			tempRest = 0
		}
	}
	if totalRest < 0 {
		return -1
	} else {
		return index
	}
}

// No.135 分糖果
func Candy(ratings []int) int {
	n := len(ratings)
	// 总消耗的糖果数
	candies := make([]int, n)
	for index := range candies {
		candies[index] = 1
	}
	// 从左到右保证分更高的有更多糖果
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// 从右到左保证分更高的有更多糖果
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i+1]+1, candies[i])
		}
	}
	total := 0
	for _, num := range candies {
		total += num
	}
	return total
}

// No.860 柠檬水找零
func LemonadeChange(bills []int) bool {
	// 用一个数组存储当前5、10美元的数量
	changes := [2]int{0, 0}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			changes[0] += 1
		} else if bills[i] == 10 {
			changes[1] += 1
			if changes[0] == 0 {
				return false
			} else {
				changes[0] -= 1
			}
		} else {
			if changes[1] > 0 && changes[0] > 0 {
				changes[1] -= 1
				changes[0] -= 1
			} else if changes[0] >= 3 {
				changes[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// No.452 用最少数量的箭引爆气球
func FindMinArrowShots(points [][]int) int {
	// 按照起始坐标升序排序，如果相同再按照终止坐标升序排序
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	total := 0
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			total++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}
	return total
}

// No.435 无重叠区间
func EraseOverlapIntervals(intervals [][]int) int {
	// 按照起始坐标升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 记录重叠区域个数
	overlap := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
			overlap++
		}
	}
	return overlap
}

// No.763 划分字母区间
func PartitionLabels(s string) []int {
	dic := [26]int{}
	for i, b := range []byte(s) {
		if i > dic[b-'a'] {
			dic[b-'a'] = i
		}
	}
	result := []int{}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, dic[s[i]-'a'])
		if i == right {
			result = append(result, right-left+1)
			left = i + 1
		}
	}
	return result
}

// No.56 合并区间
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			pre[1] = max(intervals[i][1], intervals[i-1][1])
		} else {
			ans = append(ans, pre)
			pre = intervals[i]
		}
	}
	ans = append(ans, pre)
	return ans
}

// No.738 单调递增的数字
func MonotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	str := []byte(s)
	if len(str) == 1 {
		return n
	}
	for i := len(str) - 1; i > 0; i-- {
		if str[i-1] > str[i] {
			str[i-1]--
			for j := i; j < len(str); j++ {
				str[j] = '9'
			}
		}
	}
	n, _ = strconv.Atoi(string(str))
	return n
}

// No.968 监控二叉树
func MinCameraCover(root *TreeNode) int {
	result := 0
	// 分为三个状态，0代表无覆盖，1代表有摄像头，2代表有覆盖
	var traversal func(node *TreeNode) int
	traversal = func(node *TreeNode) int {
		if node == nil {
			// 空结点说明到底了，认为是有覆盖的，因为叶子节点不放摄像头
			return 2
		}
		left := traversal(node.Left)
		right := traversal(node.Right)
		if left == 2 && right == 2 {
			// 左右孩子结点都是有覆盖，当前节点则无覆盖
			return 0
		}
		if left == 0 || right == 0 {
			// 左右孩子节点至少有一个无覆盖，当前节点需要摄像头
			result++
			return 1
		}
		if left == 1 || right == 1 {
			// 左右孩子节点至少有一个有摄像头，当前节点已覆盖
			return 2
		}
		// 并不会有这种情况，上面为了逻辑清晰没有用else
		return -1
	}
	if traversal(root) == 0 {
		result++
	}
	return result
}
