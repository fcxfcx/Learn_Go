package leetcode

import "math"

// 回文数
func IsPalindromeInt(x int) bool {
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

// 加一
func PlusOne(digits []int) []int {
	// 直接把进位开始设置为1，相当于最后一位加一
	up := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if (digit+up)/10 < 1 {
			// 不够进位
			digits[i] = digit + up
			up = 0
			return digits
		}
		// 够进位
		digits[i] = (digit + up) % 10
		up = 1
	}
	// 到最后还没有返回则说明第一位也进位了
	return append([]int{1}, digits...)
}

// 阶乘后的零
func TrailingZeroes(n int) int {
	temp, i := 1, 0
	for temp < n {
		// 计算n最多大于5的几次方
		temp *= 5
		i++
	}
	numOfFive := 0
	// 计算完全拆分后有多少个5，每有一个5就多一个0
	for j := 1; j < i; j++ {
		numOfFive += n / int(math.Pow(5, float64(j)))
	}
	return numOfFive
}

// x的平方根
func MySqrt(x int) (ans int) {
	l, r := 0, x
	for l <= r {
		// 二分法寻找
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return
}
