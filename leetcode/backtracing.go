package leetcode

import "sort"

// 电话号码的字母组合
func LetterCombinations(digits string) []string {
	words := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	if len(digits) == 0 {
		return []string{}
	}
	var dfs func(digits string) []string
	dfs = func(digits string) []string {
		result := make([]string, 0)
		if len(digits) == 1 {
			// 最后一位
			for _, word := range words[digits[0]-'2'] {
				result = append(result, string(word))
			}
		} else {
			for _, word := range words[digits[0]-'2'] {
				tempWord := string(word)
				for _, afterString := range dfs(digits[1:]) {
					result = append(result, tempWord+afterString)
				}
			}
		}
		return result
	}
	return dfs(digits)
}

// 组合
func Combine(n int, k int) (ans [][]int) {
	path := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) > k || len(path)+(n-i+1) < k {
			// 如果已经装填完了或者后面的数字不够装填
			return
		}
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for next := i; next <= n; next++ {
			path = append(path, next)
			dfs(next)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return
}

// 全排列
func Permute(nums []int) (ans [][]int) {
	hashset := make(map[int]bool, 0)
	for _, num := range nums {
		hashset[num] = true
	}
	path := []int{}
	var dfs func(length int)
	dfs = func(length int) {
		if length == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for num := range hashset {
			if hashset[num] {
				path = append(path, num)
				hashset[num] = false
				dfs(length - 1)
				path = path[:len(path)-1]
				hashset[num] = true
			}
		}
	}
	dfs(len(nums))
	return ans
}

// 组合总合
func combinationSum(candidates []int, target int) (ans [][]int) {
	path := []int{}
	// 升序排序
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var dfs func(start int)
	dfs = func(start int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			target -= candidates[i]
			dfs(i)
			target += candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
