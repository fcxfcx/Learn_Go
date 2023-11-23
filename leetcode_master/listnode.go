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
