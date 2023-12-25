package leetcode_master

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
