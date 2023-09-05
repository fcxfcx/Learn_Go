package leetcode

import "strconv"

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
