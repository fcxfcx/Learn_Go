package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

// 使用完全二叉树实现最小堆，不依靠Go语言自己的接口
type DIYMinHeap struct {
	arr []int
}

func NewMinHeap() *DIYMinHeap {
	// 构造方法
	return &DIYMinHeap{}
}

func (h *DIYMinHeap) Push(val int) {
	// 新数字加入到末尾，然后上浮
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *DIYMinHeap) Pop() int {
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

func (h *DIYMinHeap) heapifyUp(index int) {
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

func (h *DIYMinHeap) heapifyDown(index int) {
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
	hp := &DIYMinHeap{}
	for _, num := range nums {
		hp.Push(num)
	}
	for i := 0; i < len(nums)-k; i++ {
		hp.Pop()
		fmt.Println(hp)
	}
	return hp.Pop()
}

// IPO
func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
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

// 数据流的中位数
type MedianFinder struct {
	maxHeap *MaxHeap // 左边存最大堆
	minHeap *MinHeap // 右边存最小堆
	length  int
}

func MFConstructor() MedianFinder {
	return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
		length:  0,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	n := mf.length
	length_min, length_max := mf.minHeap.Len(), mf.maxHeap.Len()
	if n%2 == 0 {
		// 如果是偶数的话
		if mf.maxHeap.Len() == 0 || num < mf.maxHeap.Peek().(int) {
			// 当前数比左边最大值小，或者左边长度为零，则放左边
			heap.Push(mf.maxHeap, num)
		} else {
			// 否则放右边
			heap.Push(mf.minHeap, num)
		}
	} else {
		// 如果是奇数，则必有一边数量比另一边大一
		if length_max > length_min {
			// 左边多，则加入右边
			if num < mf.maxHeap.Peek().(int) {
				// 如果需要加入的数比左边的最大值要小，那它本来应该加入左边，因此需要调整左右长度
				// 将左边最大的加入到右边，然后将当前数字加入左边，这样左右再次均衡长度
				leftMax := heap.Pop(mf.maxHeap).(int)
				heap.Push(mf.maxHeap, num)
				heap.Push(mf.minHeap, leftMax)
			} else {
				// 否则正常加入右边
				heap.Push(mf.minHeap, num)
			}
		} else {
			// 如果右边多就加入左边，同上
			if num > mf.minHeap.Peek().(int) {
				// 如果需要加入的数比右边的最小值大，则本来应该加入右边，因此需要调整长度
				// 将右边最小的值加入到左边，然后将当前数字加入右边
				rightMin := heap.Pop(mf.minHeap).(int)
				heap.Push(mf.minHeap, num)
				heap.Push(mf.maxHeap, rightMin)
			} else {
				// 否则正常加入左边
				heap.Push(mf.maxHeap, num)
			}
		}
	}
	mf.length += 1
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() == 0 {
		return float64(mf.minHeap.Peek().(int))
	} else if mf.minHeap.Len() == 0 {
		return float64(mf.maxHeap.Peek().(int))
	}
	leftMax := mf.maxHeap.Peek().(int)
	rightMin := mf.minHeap.Peek().(int)
	if mf.length%2 == 0 {
		return float64(leftMax+rightMin) / 2.0
	} else {
		if mf.maxHeap.Len() > mf.minHeap.Len() {
			return float64(leftMax)
		} else {
			return float64(rightMin)
		}
	}
}

// 使用Go语言自带的接口实现最大堆和最小堆
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
func (hp MaxHeap) Peek() interface{} { return hp.IntSlice[0] }

type MinHeap struct{ sort.IntSlice }

func (hp MinHeap) Len() int            { return len(hp.IntSlice) }
func (hp MinHeap) Less(i, j int) bool  { return hp.IntSlice[i] < hp.IntSlice[j] }
func (hp MinHeap) Swap(i, j int)       { hp.IntSlice[i], hp.IntSlice[j] = hp.IntSlice[j], hp.IntSlice[i] }
func (hp *MinHeap) Push(x interface{}) { hp.IntSlice = append(hp.IntSlice, x.(int)) }
func (hp *MinHeap) Pop() interface{} {
	n := len(hp.IntSlice)
	result := (hp.IntSlice)[n-1]
	hp.IntSlice = hp.IntSlice[:n-1]
	return result
}
func (hp MinHeap) Peek() interface{} { return hp.IntSlice[0] }
