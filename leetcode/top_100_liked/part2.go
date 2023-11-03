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
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:length/2])
	right := mergeKLists(lists[length/2:])
	return CombineTwoList(left, right)
}
