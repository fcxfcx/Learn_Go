package leetcode

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		// 负数肯定不是回文数
		return false
	}
	num, cur := x, 0
	for num != 0 {
		cur = cur*10 + num%10
		num /= 10
	}
	return x == cur
}
