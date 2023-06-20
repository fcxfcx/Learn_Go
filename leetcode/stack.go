package leetcode

// 有效的括号
func isValidBracket(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	stack := make([]byte, 0)
	dic := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < length; i++ {
		if dic[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != dic[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
