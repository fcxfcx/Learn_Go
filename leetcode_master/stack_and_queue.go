package leetcode_master

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
