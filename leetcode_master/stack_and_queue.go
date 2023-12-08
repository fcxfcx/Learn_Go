package leetcode_master

import (
	"strconv"
)

// No.232 使用栈实现队列
type MyQueue struct {
	inStack  []int
	outStack []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	value := q.Peek()
	q.outStack = q.outStack[:len(q.outStack)-1]
	return value
}

func (q *MyQueue) Peek() int {
	lenOut := len(q.outStack)
	if lenOut == 0 {
		for i := len(q.inStack) - 1; i >= 0; i-- {
			q.outStack = append(q.outStack, q.inStack[i])
		}
		q.inStack = []int{}
	}
	value := q.outStack[len(q.outStack)-1]
	return value
}

func (q *MyQueue) Empty() bool {
	if len(q.inStack) == 0 && len(q.outStack) == 0 {
		return true
	} else {
		return false
	}
}

// No.225 用队列使用栈
type MyStack struct {
	Queue []int
}

func MyStackConstructor() MyStack {
	return MyStack{
		Queue: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	s.Queue = append(s.Queue, x)
}

func (s *MyStack) Pop() int {
	n := len(s.Queue)
	for n != 0 {
		val := s.Queue[0]
		s.Queue = s.Queue[1:]
		s.Queue = append(s.Queue, val)
		n--
	}
	res := s.Queue[0]
	s.Queue = s.Queue[1:]
	return res
}

func (s *MyStack) Top() int {
	val := s.Pop()
	s.Queue = append(s.Queue, val)
	return val
}

func (s *MyStack) Empty() bool {
	return len(s.Queue) == 0
}

// No.20 有效的括号
func IsValid(s string) bool {
	stack := []byte{}
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			stack = append(stack, ')')
		} else if str[i] == '[' {
			stack = append(stack, ']')
		} else if str[i] == '{' {
			stack = append(stack, '}')
		} else if len(stack) == 0 || stack[len(stack)-1] != str[i] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// No.1047 删除字符串中的所有相邻重复项
func RemoveDuplicates(s string) string {
	stack := []byte{}
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if len(stack) == 0 {
			stack = append(stack, str[i])
			continue
		}
		top := stack[len(stack)-1]
		if str[i] == top {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str[i])
		}
	}
	return string(stack)
}

// No.150 逆波兰表达式求值
func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if token == "+" {
				stack = append(stack, num1+num2)
			} else if token == "-" {
				stack = append(stack, num1-num2)
			} else if token == "*" {
				stack = append(stack, num1*num2)
			} else {
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

// No. 239.滑动窗口最大值
func MaxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	res := []int{}
	push := func(x int) {
		for len(q) != 0 && nums[q[len(q)-1]] < nums[x] {
			q = q[:len(q)-1]
		}
		q = append(q, x)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	for i := k - 1; i < len(nums); i++ {
		res = append(res, nums[q[0]])
		if q[0] == i-k+1 {
			q = q[1:]
		}
		push(i + 1)
	}
	return res
}
