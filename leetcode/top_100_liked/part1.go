package top_100_liked

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
