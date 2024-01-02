package leetcode_master

import (
	"sort"
	"strconv"
	"strings"
)

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
	var backtrack func(i, total int)
	backtrack = func(i, total int) {
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
			backtrack(next, total+next)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
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

// No.131 分割回文串
func Partition(s string) (res [][]string) {
	n := len(s)
	isPlalindrome := make([][]bool, n)
	for i := range isPlalindrome {
		isPlalindrome[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				isPlalindrome[i][j] = true
			} else if j-i == 1 {
				isPlalindrome[i][j] = s[i] == s[j]
			} else {
				isPlalindrome[i][j] = s[i] == s[j] && isPlalindrome[i+1][j-1]
			}
		}
	}

	path := []string{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if start >= n {
			res = append(res, append([]string(nil), path...))
			return
		}
		for end := start; end < n; end++ {
			if !isPlalindrome[start][end] {
				continue
			}
			path = append(path, s[start:end+1])
			backtrack(end + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}

// No.93 复原IP地址
func RestoreIpAddresses(s string) (res []string) {
	path := []string{}
	n := len(s)
	var backtrack func(start int)
	backtrack = func(start int) {
		if start >= n && len(path) == 4 {
			res = append(res, strings.Join(path, "."))
			return
		}
		for end := start + 1; end <= start+3 && end <= n; end++ {
			if s[start] == '0' && end > start+1 {
				// 先导0
				break
			}
			val, _ := strconv.Atoi(s[start:end])
			if val > 255 {
				continue
			}
			path = append(path, s[start:end])
			backtrack(end)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}

// No.78 子集
func Subsets(nums []int) (res [][]int) {
	path := []int{}
	sort.Ints(nums)
	var bakctrack func(start int)
	bakctrack = func(start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			bakctrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	bakctrack(0)
	return
}

// No.90 子集Ⅱ
func SubsetsWithDup(nums []int) (res [][]int) {
	path := []int{}
	sort.Ints(nums)
	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}

// No.491 非递减子序列
func FindSubsequences(nums []int) (res [][]int) {
	path := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		used := map[int]bool{}
		for i := start; i < len(nums); i++ {
			if (start != 0 && nums[i] < path[len(path)-1]) || used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}

// No.46 全排列
func Permute(nums []int) (res [][]int) {
	hashset := make(map[int]bool, 0)
	path := []int{}
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if hashset[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			hashset[nums[i]] = true
			backtrack()
			hashset[nums[i]] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return
}

// No.47 全排列2
func PermuteUnique(nums []int) (res [][]int) {
	hashset := make(map[int]bool, 0)
	path := []int{}
	sort.Ints(nums)
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && hashset[nums[i-1]] {
				continue
			}
			if !hashset[nums[i]] {
				path = append(path, nums[i])
				hashset[nums[i]] = true
				backtrack()
				hashset[nums[i]] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack()
	return
}

// No.332 重新安排行程
type pair struct {
	// pair储存机票目的地和当前航线是否已经使用
	target  string
	visited bool
}

// 储存pair数组，实现sort接口
type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

func FindItinerary(tickets [][]string) (res []string) {
	// record储存每个机场可到达的目的地，以及当前航线是否已选择
	targets := make(map[string]pairs, 0)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}
		// 添加新的可达目的地，初始置为未选择
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}
	for k := range targets {
		// 按字典升序排序目的地，保证第一个选出的就是字典序最小的结果
		sort.Sort(targets[k])
	}
	var backtrack func() bool
	backtrack = func() bool {
		if len(tickets)+1 == len(res) {
			// 结果机场数量比票数大1说明已经构建出所有路径，找到了结果
			return true
		}
		// 当前所在机场
		here := res[len(res)-1]
		for i, t := range targets[here] {
			if i > 0 && targets[here][i-1].target == t.target && !targets[here][i-1].visited {
				// 剪枝，如果上一个目的地和当前的相同，且上一个没用过，说明是从上一个回溯回来的
				// 上一个不可能那么当前的也不可能，直接跳过
				continue
			}
			// 枚举所有可能的目的地，已经使用过的航线除外
			if !t.visited {
				res = append(res, t.target)
				t.visited = true
				if backtrack() {
					return true
				}
				res = res[:len(res)-1]
				t.visited = false
			}
		}
		return false
	}
	// 所有机票从JFK出发
	res = append(res, "JFK")
	backtrack()
	return
}

// No.51 N皇后
func SolveNQueens(n int) (res [][]string) {
	// 判断某一列上有皇后
	column := make([]bool, n)
	// 判断两个对角线上有皇后
	diagonal_1, diagonal_2 := make(map[int]bool, 0), make(map[int]bool, 0)
	// 储存每行的皇后位置
	rows := make([]string, 0)
	var backtrack func(row int)
	backtrack = func(row int) {
		if len(rows) == n {
			res = append(res, append([]string{}, rows...))
			return
		}
		for i := 0; i < n; i++ {
			// 枚举当前行每一列的可能性
			if column[i] {
				continue
			}
			if diagonal_1[i+row] {
				continue
			}
			if diagonal_2[i-row] {
				continue
			}
			rows = append(rows, buildQString(i, n))
			column[i] = true
			diagonal_1[i+row] = true
			diagonal_2[i-row] = true
			backtrack(row + 1)
			rows = rows[:len(rows)-1]
			column[i] = false
			diagonal_1[i+row] = false
			diagonal_2[i-row] = false
		}
	}
	backtrack(0)
	return
}

func buildQString(i, n int) string {
	res := []byte{}
	for j := 0; j < n; j++ {
		if j != i {
			res = append(res, '.')
		} else {
			res = append(res, 'Q')
		}
	}
	return string(res)
}
