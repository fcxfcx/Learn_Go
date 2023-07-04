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

func reverseLinkList(head *ListNode) *ListNode {
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

// 反转列表Ⅱ
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right || head.Next == nil {
		return head
	}
	dummy_head := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy_head
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy_head.Next
}

// k个一组反转链表
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy_head := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy_head
	end := dummy_head
	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverseLinkList(start)
		start.Next = next
		pre = start
		end = pre
	}
	return dummy_head.Next
}

// 删除链表的倒数第 N 个结点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyhead := &ListNode{
		Val:  -1,
		Next: head,
	}
	if n == 1 && head.Next == nil {
		return nil
	}
	start, end := dummyhead, head
	count := 1
	for end.Next != nil {
		end = end.Next
		count++
	}
	for i := 0; i < count-n; i++ {
		start = start.Next
	}
	if start.Next == end {
		start.Next = nil
		return dummyhead.Next
	}
	next := start.Next.Next
	start.Next = next
	return dummyhead.Next
}
