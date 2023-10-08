package top_100_liked

import "sort"

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
