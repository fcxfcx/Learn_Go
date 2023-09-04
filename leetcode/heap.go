package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

// 使用完全二叉树实现最小堆，不依靠Go语言自己的接口
type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	// 构造方法
	return &MinHeap{}
}

func (h *MinHeap) Push(val int) {
	// 新数字加入到末尾，然后上浮
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) Pop() int {
	if len(h.arr) == 0 {
		panic("Heap is empty")
	}

	root := h.arr[0]
	last := len(h.arr) - 1

	// 首尾交换之后删除尾部
	h.arr[0] = h.arr[last]
	h.arr = h.arr[:last]

	// 交换到首部的需要向下沉
	if len(h.arr) > 0 {
		h.heapifyDown(0)
	}

	return root
}

func (h *MinHeap) heapifyUp(index int) {
	// 对index点处的数进行上浮操作
	parent := (index - 1) / 2

	for index > 0 && h.arr[index] < h.arr[parent] {
		// 当index处不满足最小堆要求的时候
		// 交换index节点和父节点
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
			// 如果当前索引的数就是最小的了，则流程停止
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

// 使用最大堆的方法
// 此处使用的是go语言自带的heap.Interface接口来实现
type MaxHeap struct{ sort.IntSlice }

func (hp MaxHeap) Len() int            { return len(hp.IntSlice) }
func (hp MaxHeap) Less(i, j int) bool  { return hp.IntSlice[i] > hp.IntSlice[j] }
func (hp MaxHeap) Swap(i, j int)       { hp.IntSlice[i], hp.IntSlice[j] = hp.IntSlice[j], hp.IntSlice[i] }
func (hp *MaxHeap) Push(x interface{}) { hp.IntSlice = append(hp.IntSlice, x.(int)) }
func (hp *MaxHeap) Pop() interface{} {
	n := len(hp.IntSlice)
	result := (hp.IntSlice)[n-1]
	hp.IntSlice = hp.IntSlice[:n-1]
	return result
}

// IPO
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	type pair struct{ c, p int }
	arr := make([]pair, n)
	for i, profit := range profits {
		arr[i] = pair{capital[i], profit}
	}
	// 将项目按照所需资金从低到高排序
	sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })
	// 使用最大堆
	hp := &MaxHeap{}
	for cur := 0; k > 0; k-- {
		for cur < n && arr[cur].c <= w {
			// 将当前资本所支持的项目加入堆中
			// 堆中按利润排序
			heap.Push(hp, arr[cur].p)
			cur++
		}
		if hp.Len() == 0 {
			// 堆为空说明无项目可做了，直接退出
			break
		}
		w += heap.Pop(hp).(int)
	}
	return w
}

// 使用最小堆
// 此处使用Go语言自带的堆接口
type tuple struct{ total, i, j int }
type TupleHeap []tuple

func (hp TupleHeap) Len() int            { return len(hp) }
func (hp TupleHeap) Swap(i, j int)       { hp[i], hp[j] = hp[j], hp[i] }
func (hp TupleHeap) Less(i, j int) bool  { return hp[i].total < hp[j].total }
func (hp *TupleHeap) Push(x interface{}) { *hp = append(*hp, x.(tuple)) }
func (hp *TupleHeap) Pop() interface{} {
	n := len(*hp)
	result := (*hp)[n-1]
	*hp = (*hp)[:n-1]
	return result
}

// 查找和最小的K对数字
func KSmallestPairs(nums1 []int, nums2 []int, k int) (result [][]int) {
	m, n := len(nums1), len(nums2)
	hp := &TupleHeap{}
	for i := 0; i < k && i < m; i++ {
		// 将(i,0)全部加入堆中
		heap.Push(hp, tuple{nums1[i] + nums2[0], i, 0})
	}
	for hp.Len() > 0 && len(result) < k {
		t := heap.Pop(hp).(tuple)
		i, j := t.i, t.j
		result = append(result, []int{nums1[i], nums2[j]})
		if j+1 < n {
			// 如果第二个数组还有数字，则加入i, j+1
			heap.Push(hp, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return
}
