package top_100_liked

func SortList(head *ListNode) *ListNode {
	return SortOne(head, nil)
}

func SortOne(head *ListNode, tail *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	// 二分法，寻找中点
	fast, slow := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != tail {
			fast = fast.Next
		}
	}
	mid := slow
	mid.Next = nil
	return CombineTwoList(SortOne(head, mid), SortOne(mid, tail))
}

func CombineTwoList(h1, h2 *ListNode) *ListNode {
	// 合并两个排序列表
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, h1, h2
	for temp1 != nil && temp2 != nil {
		if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		} else {
			temp.Next = temp1
			temp1 = temp1.Next
		}
		temp = temp.Next
	}
	// 链接剩余部分
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummy.Next
}

// 合并k个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	left := MergeKLists(lists[:length/2])
	right := MergeKLists(lists[length/2:])
	return CombineTwoList(left, right)
}

// LRU
type LRUCache struct {
	Capacity  int
	DummyHead *LRUNode
	DummyTail *LRUNode
	Hash      map[int]*LRUNode
}

type LRUNode struct {
	Next *LRUNode
	Pre  *LRUNode
	Val  int
	Key  int
}

func Constructor(capacity int) LRUCache {
	dummyHead, dummyTail := &LRUNode{}, &LRUNode{}
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return LRUCache{
		Capacity:  capacity,
		DummyHead: dummyHead,
		DummyTail: dummyTail,
		Hash:      make(map[int]*LRUNode),
	}
}

func (lc *LRUCache) DeleteHead() {
	head := lc.DummyHead.Next
	lc.DummyHead.Next = head.Next
	lc.DummyHead.Next.Pre = lc.DummyHead
	head.Next = nil
	head.Pre = nil
	delete(lc.Hash, head.Key)
}

func (lc *LRUCache) MoveToTail(node *LRUNode) {
	tempPre, tempNext := node.Pre, node.Next
	tempPre.Next = tempNext
	tempNext.Pre = tempPre
	tail := lc.DummyTail.Pre
	tail.Next = node
	node.Pre = tail
	node.Next = lc.DummyTail
	lc.DummyTail.Pre = node
}

func (lc *LRUCache) AddToTail(node *LRUNode) {
	tail := lc.DummyTail.Pre
	tail.Next = node
	node.Pre = tail
	node.Next = lc.DummyTail
	lc.DummyTail.Pre = node
}

func (lc *LRUCache) Get(key int) int {
	if node, ok := lc.Hash[key]; ok {
		// 被索引后，放置到队尾
		lc.MoveToTail(node)
		return node.Val
	} else {
		return -1
	}
}

func (lc *LRUCache) Put(key int, value int) {
	if node, ok := lc.Hash[key]; ok {
		node.Val = value
		lc.MoveToTail(node)
	} else {
		if len(lc.Hash) == lc.Capacity {
			// 超出容量，删除队头
			lc.DeleteHead()
		}
		// 新节点加入队尾
		newNode := &LRUNode{
			Val: value,
			Key: key,
		}
		lc.Hash[key] = newNode
		lc.AddToTail(newNode)
	}
}
