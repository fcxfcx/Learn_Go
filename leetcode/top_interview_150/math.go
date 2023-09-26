package top_interview_150

import (
	"math"
	"strconv"
)

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

// Pow(x,n)
func MyPow(x float64, n int) float64 {
	var quickMul func(x float64, n int) float64
	quickMul = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		y := quickMul(x, n/2)
		if n%2 == 0 {
			// 可以拆分成相同的两部分
			return y * y
		}
		return y * y * x
	}
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}

// 直线上最多的点数
func MaxPoints(points [][]int) (ans int) {
	n := len(points)
	// 特判
	if n < 3 {
		// 点小于三个则直接返回点的数量
		return n
	}
	for i, p := range points {
		if ans > n-i || ans > n/2 {
			// 遍历到第i个点的时候，最多就能找到n-i个点在线上，因为前面都遍历过了
			// 因此如果找到的最大值已经大于这个值了，就没必要再找了
			// 另外，如果找到的最大值已经比一半的点多了，就不用再找了
			break
		}
		count := map[string]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					// 保证如果有负数只让分子为负数，避免哈希表重复记录同一条线
					x, y = -x, -y
				}
				// 求最大公约数，将直线表示为最简形式
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			// 斜率相同则在一条直线上
			key := strconv.Itoa(x) + "@" + strconv.Itoa(y)
			count[key]++
		}
		for _, c := range count {
			ans = max(ans, c+1)
		}
	}
	return
}

func gcd(a, b int) int {
	// 辗转相除法求最大公约数
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
