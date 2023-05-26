package leetcode

import (
	"fmt"
	"math"
	"math/rand"
)

// 合并两个有序数组
func Merge(nums1 []int, m int, nums2 []int, n int) {
	m_index, n_index, cur := m-1, n-1, 0
	for i := m + n - 1; i >= 0; i-- {
		if m_index == -1 {
			cur = nums2[n_index]
			n_index--
		} else if n_index == -1 {
			cur = nums1[m_index]
			m_index--
		} else if nums1[m_index] < nums2[n_index] {
			cur = nums2[n_index]
			n_index--
		} else {
			cur = nums1[m_index]
			m_index--
		}
		nums1[i] = cur
	}
	fmt.Println(nums1)
}

// 移除数组中某一元素
func RemoveElement(nums []int, val int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[head] == val {
			nums[head] = nums[tail]
			tail--
		} else {
			head++
		}
	}
	fmt.Println(nums)
	return head
}

// 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	head, tail := 0, 1
	for tail < len(nums) {
		if nums[head] == nums[tail] {
			tail++
		} else {
			nums[head+1] = nums[tail]
			head++
		}
	}
	fmt.Println(nums)
	return head + 1
}

// 删除有序数组中的重复项Ⅱ 重复项最多可保留两个
func RemoveDuplicatesPlus(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow := 2
	for fast := 2; fast < n; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 返回数组中数量大于一半的元素
func MajorityElement(nums []int) int {
	count, cur := 0, 0
	for _, num := range nums {
		if count == 0 {
			cur = num
			count++
		} else if cur != num {
			count--
		} else {
			count++
		}
	}
	return cur
}

// 轮转数组元素
func Rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// 买卖股票的最佳时机
func MaxProfit(prices []int) int {
	min, profit := math.MaxInt64, 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		profit = int(math.Max(float64(profit), float64(price-min)))
	}
	return profit
}

// 买卖股票的最佳时机Ⅱ（多次买卖）
func MaxProfitPlus(prices []int) int {
	length := len(prices)
	dp := make([][2]int, length)
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}

// 跳跃游戏
func CanJump(nums []int) bool {
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if i > maxLength {
			return false
		} else {
			maxLength = max(maxLength, i+nums[i])
		}
	}
	return true
}

// H指数
func HIndex(citations []int) int {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	for i, total := n, 0; i > 0; i-- {
		total += counter[i]
		if total >= i {
			return i
		}
	}
	return 0
}

// O(1) 时间插入、删除和获取随机元素
type RandomizedSet struct {
	hashmap  map[int]int
	elements []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashmap:  make(map[int]int),
		elements: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.hashmap[val]; ok {
		return false
	}
	rs.hashmap[val] = len(rs.elements)
	rs.elements = append(rs.elements, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, ok := rs.hashmap[val]
	if !ok {
		return false
	}
	last := len(rs.elements) - 1
	rs.elements[index] = rs.elements[last]
	rs.hashmap[rs.elements[last]] = index
	rs.elements = rs.elements[:last]
	delete(rs.hashmap, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.elements[rand.Intn(len(rs.elements))]
}

// 除自身以外数组的乘积
func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	answer, L, R := make([]int, length), make([]int, length), make([]int, length)
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

// 加油站
func CanCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumGas, sumCost, canCover := 0, 0, 0
		for canCover < n {
			j := (i + canCover) % n
			sumGas += gas[j]
			sumCost += cost[j]
			if sumGas < sumCost {
				break
			}
			canCover++
		}
		if canCover == n {
			return i
		} else {
			i += canCover + 1
		}
	}
	return -1
}

// 接雨水
func Trap(height []int) int {
	length := len(height)
	left, right, result := 0, length-1, 0
	leftMaxHeight := height[left]
	rightMaxHeight := height[right]
	for left < right {
		if leftMaxHeight < rightMaxHeight {
			result += leftMaxHeight - height[left]
			if left++; height[left] > leftMaxHeight {
				leftMaxHeight = height[left]
			}
		} else {
			result += rightMaxHeight - height[right]
			if right--; height[right] > rightMaxHeight {
				rightMaxHeight = height[right]
			}
		}
	}
	return result
}

// 分糖果
func Candy(ratings []int) (ans int) {
	length := len(ratings)
	left, right := make([]int, length), 0
	left[0] = 1
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for j := length - 1; j >= 0; j-- {
		if j < length-1 && ratings[j] > ratings[j+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[j], right)
	}
	return
}

// 罗马数字转整数
func RomanToInt(s string) int {
	valueMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	length, answer := len(s), 0
	for n := range s {
		value := valueMap[s[n]]
		if n < length-1 && value < valueMap[s[n+1]] {
			answer -= value
		} else {
			answer += value
		}
	}
	return answer
}

// 整数转罗马数字
func IntToRoman(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	for i := 0; i < len(value); i++ {
		for num >= value[i] {
			num -= value[i]
			roman += symbol[i]
		}
		if num == 0 {
			break
		}
	}
	return roman
}

// 最后一个单词的长度
func LengthOfLastWord(s string) int {
	temp := []byte(s)
	length, count := len(temp), 0
	for i := length - 1; i > 0; i-- {
		if temp[i] != ' ' {
			count++
		} else if count != 0 {
			return count
		}
	}
	return count
}

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	result := ""
	for i := 0; i < len(strs[0]); i++ {
		tempByte := strs[0][i]
		for _, s := range strs {
			if i > len(s) || s[i] != tempByte {
				return result
			}
		}
		result += string(tempByte)
	}
	return result
}

// ----------- 私有工具类方法 --------------
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
