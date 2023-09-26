package top_interview_150

import (
	"container/heap"
)

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

// 将有序数组转换为二叉搜索树
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	midNode := &TreeNode{
		Val: nums[mid],
	}
	midNode.Left = SortedArrayToBST(nums[0:mid])
	midNode.Right = SortedArrayToBST(nums[mid+1:])
	return midNode
}

// 排序列表
func SortList(head *ListNode) *ListNode {
	// 使用分治算法，归并排序
	return sortOne(head, nil)
}

func combineListNode(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, tempLeft, tempRight := dummy, left, right
	for tempLeft != nil && tempRight != nil {
		if tempLeft.Val <= tempRight.Val {
			temp.Next = tempLeft
			tempLeft = tempLeft.Next
		} else {
			temp.Next = tempRight
			tempRight = tempRight.Next
		}
		temp = temp.Next
	}
	if tempLeft != nil {
		temp.Next = tempLeft
	} else if tempRight != nil {
		temp.Next = tempRight
	}
	return dummy.Next
}

func sortOne(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		// 快慢指针找链表中间节点
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return combineListNode(sortOne(head, mid), sortOne(mid, tail))
}

// 建立四叉树
func Construct(grid [][]int) *QuadNode {
	var dfs func(rows [][]int, c0 int, c1 int) *QuadNode
	dfs = func(rows [][]int, c0, c1 int) *QuadNode {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] { // 不是叶节点
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &QuadNode{
						true,
						false,
						dfs(rows[:rMid], c0, cMid),
						dfs(rows[:rMid], cMid, c1),
						dfs(rows[rMid:], c0, cMid),
						dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		// 叶子节点
		return &QuadNode{Val: rows[0][c0] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}

// 合并k个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	hp := &ListNodeHeap{}
	// 建立堆
	heap.Init(hp)
	for _, list := range lists {
		// 将头结点加入堆中
		if list != nil {
			heap.Push(hp, list)
		}
	}

	dummy := &ListNode{}
	cur := dummy
	for hp.Len() > 0 {
		node := heap.Pop(hp).(*ListNode)
		if node.Next != nil {
			// 下一节点不为空则填入堆
			heap.Push(hp, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

// 使用最小堆的方法
// 此处使用的是go语言自带的heap.Interface接口来实现
type ListNodeHeap []*ListNode

func (hp ListNodeHeap) Len() int            { return len(hp) }
func (hp ListNodeHeap) Less(i, j int) bool  { return hp[i].Val < hp[j].Val }
func (hp ListNodeHeap) Swap(i, j int)       { hp[i], hp[j] = hp[j], hp[i] }
func (hp *ListNodeHeap) Push(x interface{}) { *hp = append(*hp, x.(*ListNode)) }
func (hp *ListNodeHeap) Pop() interface{} {
	n := len(*hp)
	result := (*hp)[n-1]
	*hp = (*hp)[:n-1]
	return result
}
