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
