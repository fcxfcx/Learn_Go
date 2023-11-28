package leetcode_master

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
