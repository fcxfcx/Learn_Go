package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
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
