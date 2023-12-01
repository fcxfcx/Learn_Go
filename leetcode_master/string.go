package leetcode_master

// No.344 反转字符串
func ReverseString(s []byte) {
	n := len(s)
	left, right := 0, n-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// No.541 反转字符串Ⅱ
func ReverseStr(s string, k int) string {
	str := []byte(s)
	count := len(s) / (2 * k)
	for i := 0; i < count; i++ {
		start := i * k
		end := start + k
		if end >= len(s) {
			return string(str)
		}
		for start < end {
			str[start], str[end] = str[end], str[start]
			start++
			end--
		}
	}
	return string(str)
}
