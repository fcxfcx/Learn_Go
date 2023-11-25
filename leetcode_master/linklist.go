package leetcode_master

type ListNode struct {
	Val  int
	Next *ListNode
}

// No.203 移除链表元素
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	for temp := dummy; temp.Next != nil; {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return dummy.Next
}

// No.707 设计链表(实现)
type MyLinkedList struct {
	DummyHead *ListNode
	Len       int
}

func Constructor() MyLinkedList {
	dummyHead := &ListNode{}
	return MyLinkedList{
		DummyHead: dummyHead,
		Len:       0,
	}
}

func (list *MyLinkedList) Get(index int) int {
	if index >= list.Len {
		return -1
	}
	temp := list.DummyHead
	for i := 0; i <= index; i++ {
		temp = temp.Next
	}
	return temp.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	tempHead := list.DummyHead.Next
	newHead := &ListNode{
		Val:  val,
		Next: tempHead,
	}
	list.DummyHead.Next = newHead
	list.Len += 1
}

func (list *MyLinkedList) AddAtTail(val int) {
	tempTail := list.DummyHead
	for tempTail.Next != nil {
		tempTail = tempTail.Next
	}
	newTail := &ListNode{
		Val: val,
	}
	tempTail.Next = newTail
	list.Len += 1
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index > list.Len {
		return
	}
	cur := list.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	next := cur.Next
	newNode := &ListNode{
		Val:  val,
		Next: next,
	}
	cur.Next = newNode
	list.Len += 1
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index > list.Len {
		return
	}
	cur := list.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	next := cur.Next
	cur.Next = next.Next
	list.Len -= 1
}

// No.206 反转链表
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// No.24 两两交换链表中的节点
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node1, node2, node3 := head, head.Next, head.Next.Next
	node2.Next = node1
	node1.Next = SwapPairs(node3)
	return node2
}

// No.19 删除链表的倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 面试题 02.07. 链表相交
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
