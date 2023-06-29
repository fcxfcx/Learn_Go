package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 环形链表
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := result
	plus_signal := 0
	for l1 != nil || l2 != nil {
		value := 0
		if l1 == nil {
			value = l2.Val + plus_signal
			if value >= 10 {
				plus_signal = 1
				value = value - 10
			} else {
				plus_signal = 0
			}
			cur.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			l2 = l2.Next
		} else if l2 == nil {
			value = l1.Val + plus_signal
			if value >= 10 {
				plus_signal = 1
				value = value - 10
			} else {
				plus_signal = 0
			}
			cur.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			value = l1.Val + l2.Val + plus_signal
			if value >= 10 {
				plus_signal = 1
				value = value - 10
			} else {
				plus_signal = 0
			}
			cur.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			l1 = l1.Next
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if plus_signal == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return result.Next
}

// 合并两个有序链表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		} else if list2 == nil {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			if list1.Val > list2.Val {
				cur.Next = list2
				list2 = list2.Next
				cur = cur.Next
			} else {
				cur.Next = list1
				list1 = list1.Next
				cur = cur.Next
			}
		}
	}
	return head.Next
}

// 复制带随机指针的链表
func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{
			Val:    node.Val,
			Next:   node.Next,
			Random: nil,
		}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	new_head := head.Next
	for node := head; node != nil; node = node.Next {
		newnode := node.Next
		node.Next = node.Next.Next
		if newnode.Next != nil {
			newnode.Next = newnode.Next.Next
		}
	}
	return new_head
}
