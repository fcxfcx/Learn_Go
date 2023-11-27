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
