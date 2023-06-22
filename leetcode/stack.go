package leetcode

import (
	"strings"
)

// 有效的括号
func IsValidBracket(s string) bool {
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

// 简化路径
func SimplifyPath(path string) string {
	stack := make([]string, 0)
	for _, item := range strings.Split(path, "/") {
		if item == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if item != "" && item != "." {
			stack = append(stack, item)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// 最小栈
type MinStack struct {
	stack    []int
	minStack []int
}

func StackConstructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.minStack) == 0 {
		ms.minStack = append(ms.minStack, val)
	} else if ms.GetMin() < val {
		ms.minStack = append(ms.minStack, ms.GetMin())
	} else {
		ms.minStack = append(ms.minStack, val)
	}

}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}
