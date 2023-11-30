package leetcode_master

import (
	"sort"
)

// No.242 有效的字母异位词
func IsAnagram(s string, t string) bool {
	dic := [26]int{}
	for _, b := range s {
		index := int(b - 'a')
		dic[index] += 1
	}
	for _, b := range t {
		index := int(b - 'a')
		dic[index] -= 1
	}
	for _, val := range dic {
		if val != 0 {
			return false
		}
	}
	return true
}

// No.349 两个数组的交集
func Intersection(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	res := []int{}
	for _, num1 := range nums1 {
		hash[num1] += 1
	}
	for _, num2 := range nums2 {
		if _, ok := hash[num2]; ok {
			res = append(res, num2)
			delete(hash, num2)
		}
	}
	return res
}

// No.202 快乐数
func IsHappy(n int) bool {
	hashset := map[int]bool{}
	for !hashset[n] {
		temp, new := n, 0
		hashset[n] = true
		for temp != 0 {
			new += (temp % 10) * (temp % 10)
			temp = temp / 10
		}
		if new == 1 {
			return true
		}
		n = new
	}
	return false
}

// No.1 两数之和
func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	res := [2]int{}
	for index, num := range nums {
		hash[target-num] = index
	}
	for index, num := range nums {
		if index2, ok := hash[num]; ok && index2 != index {
			res[0] = index2
			res[1] = index
		}
	}
	return res[:]
}

// No.454 四数相加Ⅱ
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	hash := map[int]int{}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			hash[nums1[i]+nums2[j]] += 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if c, ok := hash[-nums3[i]-nums4[j]]; ok {
				count += c
			}
		}
	}
	return count
}

// No.383 赎金信
func CanConstruct(ransomNote string, magazine string) bool {
	hash := [26]int{}
	for _, b := range []byte(magazine) {
		hash[b-'a'] += 1
	}
	for _, v := range []byte(ransomNote) {
		if hash[v-'a'] > 0 {
			hash[v-'a'] -= 1
		} else {
			return false
		}
	}
	return true
}

// No.15 三数之和
func ThreeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	// 升序排序
	sort.Ints(nums)
	n := len(nums)
	// 寻找a + b + c = 0
	// 假设a = nums[i], b = nums[left], c = nums[right]
	for i := 0; i < n-2; i++ {
		// 至少要留三个数，因此i需要小于n-2
		// a 的去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		if a > 0 {
			// 第一个数大于0则之后都不会有相加等于0的三个结果
			break
		}
		left, right := i+1, n-1
		for left < right {
			b, c := nums[left], nums[right]
			if a+b+c == 0 {
				ans = append(ans, []int{a, b, c})
				// b的去重
				for left < right && nums[left] == b {
					left++
				}
				// c的去重
				for left < right && nums[right] == c {
					right--
				}
			} else if a+b+c > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

// No.18 四数之和
func FourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		// 同样需要去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		for j := i + 1; j < n-2; j++ {
			if j > 1 && nums[j] == nums[j-1] {
				continue
			}
			b := nums[j]
			left, right := j+1, n-1
			for left < right {
				c, d := nums[left], nums[right]
				if a+b+c+d == target {
					ans = append(ans, []int{a, b, c, d})
					// b的去重
					for left < right && nums[left] == c {
						left++
					}
					// c的去重
					for left < right && nums[right] == d {
						right--
					}
				} else if a+b+c+d < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}
