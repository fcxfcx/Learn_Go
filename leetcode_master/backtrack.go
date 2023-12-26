package leetcode_master

import "sort"

// No.77 组合
func Combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(i int)
	backtrack = func(i int) {
		if len(path) > k || len(path)+(n-i+1) < k {
			return
		}
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for next := i + 1; next <= n; i++ {
			path = append(path, next)
			backtrack(next)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// No.216 组合总和 III
func CombinationSum3(k int, n int) (res [][]int) {
	path := []int{}
	var backtarck func(i, total int)
	backtarck = func(i, total int) {
		if len(path) > k || (n-total) > 9*(k-len(path)) {
			return
		}
		if len(path) == k && total == n {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for next := i + 1; next <= 9; next++ {
			path = append(path, next)
			backtarck(next, total+next)
			path = path[:len(path)-1]
		}
	}
	backtarck(0, 0)
	return
}

// No.17 电话号码的字母组合
func LetterCombinations(digits string) []string {
	words := [][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'},
		{'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'},
	}
	res, path := []string{}, []byte{}
	if len(digits) == 0 {
		return res
	}
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits)-1 {
			res = append(res, string(path))
			return
		}
		wordsIndex := digits[index] - '2'
		for _, b := range words[wordsIndex] {
			path = append(path, b)
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// No.39 组合总和
func CombinationSum(candidates []int, target int) (res [][]int) {
	path := []int{}
	sort.Ints(candidates)
	var backtrack func(index int)
	backtrack = func(index int) {
		if target < 0 {
			return
		}
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := index; i < len(candidates); i++ {
			val := candidates[i]
			path = append(path, val)
			target -= val
			backtrack(i)
			path = path[:len(path)-1]
			target += val
		}
	}
	backtrack(0)
	return
}

// No.40 组合总和Ⅱ
func CombinationSum2(candidates []int, target int) (res [][]int) {
	path := []int{}
	sort.Ints(candidates)
	used := make([]int, len(candidates))
	var backtrack func(index int)
	backtrack = func(index int) {
		if target < 0 {
			return
		}
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == 0 {
				// 对于同一层的遍历，需要去重
				continue
			}
			val := candidates[i]
			path = append(path, val)
			used[i] = 1
			target -= val
			backtrack(i + 1)
			used[i] = 0
			path = path[:len(path)-1]
			target += val
		}
	}
	backtrack(0)
	return
}
