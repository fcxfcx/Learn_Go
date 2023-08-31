package leetcode

import "fmt"

// 使用完全二叉树实现最小堆，不依靠Go语言自己的接口
type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Push(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) Pop() int {
	if len(h.arr) == 0 {
		panic("Heap is empty")
	}

	root := h.arr[0]
	last := len(h.arr) - 1

	h.arr[0] = h.arr[last]
	h.arr = h.arr[:last]

	if len(h.arr) > 0 {
		h.heapifyDown(0)
	}

	return root
}

func (h *MinHeap) heapifyUp(index int) {
	parent := (index - 1) / 2

	for index > 0 && h.arr[index] < h.arr[parent] {
		h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *MinHeap) heapifyDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < len(h.arr) && h.arr[left] < h.arr[smallest] {
			smallest = left
		}
		if right < len(h.arr) && h.arr[right] < h.arr[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}

		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
		index = smallest
	}
}

// 数组中第k个最大的元素
func FindKthLargest(nums []int, k int) int {
	hp := &MinHeap{}
	for _, num := range nums {
		hp.Push(num)
	}
	for i := 0; i < len(nums)-k; i++ {
		hp.Pop()
		fmt.Println(hp)
	}
	return hp.Pop()
}
