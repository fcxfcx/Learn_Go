package leetcode

import "strconv"

// 二进制求和
func AddBinary(a string, b string) string {
	m, n := len(a), len(b)
	result := make([]int, 0)
	stringAns := make([]byte, 0)
	up := 0
	for i := 0; i < max(m, n); i++ {
		temp := up
		if i < n {
			b_int, _ := strconv.Atoi(string(b[n-i-1]))
			temp += b_int
		}
		if i < m {
			a_int, _ := strconv.Atoi(string(a[m-i-1]))
			temp += a_int
		}
		if temp >= 2 {
			up = 1
			result = append(result, temp-2)
		} else {
			up = 0
			result = append(result, temp)
		}
	}
	if up == 1 {
		result = append(result, 1)
	}
	for j := len(result) - 1; j >= 0; j -= 1 {
		stringAns = append(stringAns, byte('0'+result[j]))
	}
	return string(stringAns)
}

// 颠倒二进制位
func ReverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		res = (res << 1) | (num & 1)
		num = num >> 1
	}
	return res
}

// 位1的个数
func HammingWeight(num uint32) (count int) {
	for ; num > 0; num &= (num - 1) {
		count++
	}
	return
}

// 只出现一次的数字
func SingleNumber(nums []int) (res int) {
	// 交换律：a ^ b ^ c <=> a ^ c ^ b
	// 任何数于0异或为任何数 0 ^ n => n
	// 相同的数异或为0: n ^ n => 0
	for _, num := range nums {
		res = res ^ num
	}
	return
}

// 只出现一次的数字Ⅱ
func SingleNumberPlus(nums []int) int {
	// 对于数组中非答案的元素，每一个元素都出现了3次，对应着第i个二进制位的3个0或3个1
	// 无论是哪一种情况，它们的和都是3的倍数（即和为 0 或 3）
	// 答案的第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数。
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

// 数字范围按位与
func RangeBitwiseAnd(left int, right int) int {
	// 对所有数字执行按位与运算的结果是所有对应二进制字符串的公共前缀再用零补上后面的剩余位
	// 所有这些二进制字符串的公共前缀也即指定范围的起始和结束数字left和right的公共前缀
	// 因此就找left和right的公共前缀
	shift := 0
	for left < right {
		left, right = left>>1, right>>1
		shift++
	}
	return left << shift
}

