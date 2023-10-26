package top_100_liked

import (
	"math"
	"sort"
)

// 两数之和
func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for i, num := range nums {
		hash[target-num] = i
	}
	for j := range nums {
		if i, ok := hash[nums[j]]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

// 字母异位词分组
func GroupAnagrams(strs []string) (res [][]string) {
	hash := make(map[[26]int][]string)
	for _, str := range strs {
		temp := [26]int{}
		for _, c := range str {
			temp[c-'a'] += 1
		}
		hash[temp] = append(hash[temp], str)
	}
	for i := range hash {
		res = append(res, hash[i])
	}
	return
}

// 最长连续序列
func LongestConsecutive(nums []int) int {
	// 存一个hashmap，键是序列首尾组合，值是连续序列长度
	hash := make(map[int]bool)
	maxLength := 0
	for _, num := range nums {
		hash[num] = true
	}
	for _, num := range nums {
		if !hash[num-1] {
			// 从序列起点开始遍历，如果map中有num-1，那之后也会遍历到的
			curNum, curLength := num, 1
			for hash[curNum+1] {
				curNum++
				curLength++
			}
			if curLength > maxLength {
				maxLength = curLength
			}
		}
	}
	return maxLength
}

// 移动零
func MoveZeroes(nums []int) {
	n := len(nums)
	// a代表当前需要填数字的下标，b代表当前处理的非零值下标
	a, b := 0, 0
	for b < n {
		if nums[b] == 0 {
			// 零值就跳过
			b++
		} else {
			nums[a] = nums[b]
			a++
			b++
		}
	}
	for a < n {
		nums[a] = 0
		a++
	}
}

// 盛水最多的容器
func MaxArea(height []int) int {
	n := len(height)
	left, right, res := 0, n-1, 0
	for left < right {
		area := 0
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		res = max(res, area)
	}
	return res
}

// 三数之和
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	// 假设三个数字是a,b,c
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			// a和上一个相同则跳过
			continue
		}
		second, third := first+1, n-1
		target := -nums[first]
		for second < third {
			if nums[second]+nums[third] < target || (second > first+1 && nums[second] == nums[second-1]) {
				// b+c < -a则说明b+c不够大，需要挪动b
				// 或者b和上一个b相同也不可以
				second++
			} else if nums[second]+nums[third] > target || (third < n-1 && nums[third] == nums[third+1]) {
				// 同理如果b+c > -a 则shuom b+c不够小，需要挪动c
				// 或者c和上一个c相同也不可以
				third--
			} else if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				second++
			}
		}
	}
	return res
}

// 接雨水
func Trap(height []int) int {
	n := len(height)
	sum, left, right := 0, 0, n-1
	leftMax, rightMax := height[0], height[n-1]
	for left < right {
		if leftMax < rightMax {
			// 以左边的最大值为准
			sum += (leftMax - height[left])
			left++
			if height[left] > leftMax {
				// 移动左边指针，维护左侧最大值
				leftMax = height[left]
			}
		} else {
			// 以右边的最大值为准
			sum += (rightMax - height[right])
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			}
		}
	}
	return sum
}

// 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	hash := make(map[byte]bool, 0)
	start, end, maxLength := 0, 0, 1
	for end < n {
		if !hash[byte(s[end])] || start == end {
			hash[byte(s[end])] = true
			end++
			if end-start > maxLength {
				maxLength = end - start
			}
		} else {
			delete(hash, byte(s[start]))
			start++
		}
	}
	return maxLength
}

// 找到字符串中所有字母异位词
func FindAnagrams(s string, p string) []int {
	len_s, len_p := len(s), len(p)
	ans := []int{}
	if len_p > len_s {
		// 特判
		return ans
	}
	count := [26]int{}
	for i, s_byte := range p {
		// 将p中的字符串记录在哈希表中(用数组存储)
		count[s_byte-'a'] -= 1
		count[s[i]-'a'] += 1
	}

	// 用differ代表当前s的滑动窗口和p的不同字符数量
	differ := 0
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			differ++
		}
	}
	if differ == 0 {
		// 如果第一个滑动窗口就符合，则将0加入ans
		ans = append(ans, 0)
	}

	// 滑动窗口
	for i, ch := range s[:len_s-len_p] {
		if count[ch-'a'] == 1 {
			// 移动后，窗口内字符不同处减一
			differ--
		} else if count[ch-'a'] == 0 {
			// 否则如果这个字符已经符合过了，减掉它会加入一个不同
			differ++
		}
		count[ch-'a']--

		if count[s[i+len_p]-'a'] == -1 {
			// 窗口尾部移动后的字符为-1代表p中有，所以右边滑动后不同会变少
			differ--
		} else if count[s[i+len_p]-'a'] == 0 {
			differ++
		}
		count[s[i+len_p]-'a']++

		if differ == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// 和为k的子数组
func SubarraySum(nums []int, k int) int {
	hash := map[int]int{}
	hash[0] = 1
	pre, count := 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := hash[pre-k]; ok {
			count += hash[pre-k]
		}
		hash[pre] += 1
	}
	return count
}

// 滑动窗口最大值
func MaxSlidingWindow(nums []int, k int) []int {
	// 构造单调队列，储存数组下标
	q := make([]int, 0)
	push := func(i int) {
		for len(q) > 0 && q[len(q)-1] < i {
			// 如果新入队的数大于队尾的数，则队尾可以去除
			// 因为只要新加入的数还存在，那么队尾的这个数就不可能被选到
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	// 注意队列里存的是数组下标，但是结果要求返回的是数
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			// 队头的元素超过滑动窗口大小的剔除
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

// 最小覆盖子串
func MinWindow(s string, t string) string {
	len_s, len_t := len(s), len(t)
	if len_s < len_t {
		return ""
	}
	t_map := make(map[byte]int, 0)
	// 初始化待匹配字符串对应的哈希表
	for i := 0; i < len_t; i++ {
		t_map[byte(t[i])] += 1
	}
	// 代表结果子串和已匹配的字符数量
	result, count := "", 0
	s_map := make(map[byte]int, 0)
	for left, right := 0, 0; left <= right && right < len_s; right++ {
		temp := byte(s[right])
		// 如果这个字符在t中
		if num, ok := t_map[temp]; ok {
			s_map[temp] += 1
			if s_map[temp] <= num {
				// 待匹配的字符数量还没有超过t中的数量
				// 代表新加入的这个字符对匹配有帮助
				count++
			}
		}
		if count == len_t {
			// 如果全部匹配成功
			for left < right {
				// 右移左边界直至无法匹配
				if num, ok := s_map[byte(s[left])]; ok {
					if num == t_map[byte(s[left])] {
						// 代表当前的字符是不可去的，删去会无法匹配
						break
					} else {
						// 否则就可以右移左边界
						s_map[byte(s[left])] -= 1
						left++
					}
				} else {
					// 如果s_map里不含这个字符说明它不是t中的字符
					left++
				}
			}
			// 维护最小子串
			if len(result) == 0 || right-left < len(result) {
				result = s[left : right+1]
			}
			// 此时左边第一个一定是已批配字符，右移左边界开始新一次的匹配
			s_map[s[left]] -= 1
			left += 1
			count--
		}
	}
	return result
}

// 最大子数组和
func MaxSubArray(nums []int) int {
	sum := 0
	ans := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum = nums[i] + max(sum, 0)
		ans = max(ans, sum)
	}
	return ans
}

// 合并区间
func Merge(intervals [][]int) [][]int {
	// 数组排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	pre := intervals[0]
	for _, cur := range intervals {
		if pre[1] < cur[0] {
			// 没有重叠部分
			ans = append(ans, pre)
			pre = cur
			continue
		} else if pre[1] < cur[1] {
			pre[1] = cur[1]
		}
	}
	ans = append(ans, pre)
	return ans
}

// 轮转数组
func Rotate(nums []int, k int) {
	reverseArr := func(arr []int) {
		for i, n := 0, len(arr); i < n/2; i++ {
			arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
		}
	}
	k %= len(nums)
	reverseArr(nums)
	reverseArr(nums[:k])
	reverseArr(nums[k:])
}

// 除自身以外数组的乘积
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 首先计算每一个下标处，左边所有数的积
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// 从右往左计算每一个下标处，右边所有数的积
	R := 1
	for i := n - 1; i >= 0; i++ {
		answer[i] = answer[i] * R
		R = R * nums[i]
	}
	return answer
}

// 缺失的第一个正数
func FirstMissingPositive(nums []int) int {
	// 将数组当成原地的哈希表，在下标i处储存数字i+1
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 负数的情况，储存数字大于n的情况（最大就是n）和已经储存正确数字的情况可以跳过
			// 否则就替换当前的数
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// 矩阵置零
func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 储存第一行和第一列是否原来就含有零
	column, row := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			// 发现第一列中有一个零，则第一列记录为全零
			column = true
			break
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			// 发现第一行中有一个零，则第一行记录为全零
			row = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if column {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// 螺旋矩阵
func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	actions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	actionIndex, row, col := 0, 0, 0
	result := make([]int, m*n)
	// 特判，如果行或列等于1
	if m == 1 {
		return matrix[0]
	} else if n == 1 {
		for i, value := range matrix {
			result[i] = value[0]
		}
		return result
	}
	for i := 0; i < len(result); i++ {
		result[i] = matrix[row][col]
		// 矩阵中数字是-100至100，用-101标注已填入
		matrix[row][col] = -101
		// 遇到其他三个顶角需要拐弯
		if (row == m-1 && col == 0) || (row == m-1 && col == n-1) || (row == 0 && col == n-1) {
			actionIndex = (actionIndex + 1) % 4
		} else {
			maybeNextRow := row + actions[actionIndex][0]
			maybeNextCol := col + actions[actionIndex][1]
			if matrix[maybeNextRow][maybeNextCol] == -101 {
				actionIndex = (actionIndex + 1) % 4
			}
		}
		row += actions[actionIndex][0]
		col += actions[actionIndex][1]
	}
	return result
}

// 旋转矩阵
func RotateMatrix(matrix [][]int) {
	n := len(matrix)
	// 用两次翻转代替旋转
	// 水平上下翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 搜索二维矩阵Ⅱ
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	up, down := 0, m-1
	for up < down-1 {
		// 二分法
		middle := (down + up) / 2
		if matrix[middle][0] > target {
			// 中间的数大于target，左移右边界
			down = middle
		} else if matrix[middle][0] == target {
			return true
		} else {
			// 中间的数大于target，则target不可能在这一行
			up = middle
		}
	}
	for i := up; i >= 0; i-- {
		// 遍历行，在每一行用二分法搜寻数字
		left, right := 0, n
		for left < right {
			middle := (left + right) / 2
			if matrix[i][middle] == target {
				return true
			} else if matrix[i][middle] > target {
				right = middle
			} else {
				left = middle + 1
			}
		}
	}
	return false
}

// 相交链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
