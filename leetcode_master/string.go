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
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			ReverseString(str[i : i+k])
		} else {
			ReverseString(str[i:len(s)])
		}
	}
	return string(str)
}

// No.151 反转字符串中的单词
func ReverseWords(s string) string {
	str := []byte(s)
	n := len(s)
	// 双指针法去除多余空格
	slow, fast := 0, 0
	// 去除头部多余空格
	for fast < n && str[fast] == ' ' {
		fast++
	}
	// 去除中间多余空格
	for ; fast < n; fast++ {
		if fast > 1 && str[fast-1] == str[fast] && str[fast] == ' ' {
			// 如果出现连续空格，右移快指针，这里判定的是出现两个及以上的空格，只记录第一个空格给slow
			continue
		}
		str[slow] = str[fast]
		slow++
	}
	// 删除尾部空格
	if slow > 1 && str[slow-1] == ' ' {
		str = str[:slow-1]
	} else {
		str = str[:slow]
	}

	// 反转整个字符串
	ReverseString(str)

	// 反转每个单词
	i := 0
	for i < len(str) {
		j := i
		for ; j < len(str) && str[j] != ' '; j++ {
		}
		ReverseString(str[i:j])
		// j此时对应的是空格，所以还要加一
		i = j + 1
	}
	return string(str)
}

// No.28 找出字符串中第一个匹配项的下标
func StrStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if m < n {
		return -1
	}
	next := make([]int, n)
	makeNext(next, needle)

	j := 0
	for i := 0; i < m; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func makeNext(next []int, s string) {
	// KMP算法构造next数组
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

// No.459 重复的子字符串
func RepeatedSubstringPattern(s string) bool {
	n := len(s)
	next := make([]int, n)
	makeNext(next, s)
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}
