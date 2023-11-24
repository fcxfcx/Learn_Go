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
	stack := []*ListNode{}
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	dummy := &ListNode{}
	cur := dummy
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		cur.Next = temp
		cur = cur.Next
		stack = stack[:len(stack)-1]
	}
	return dummy.Next
}
