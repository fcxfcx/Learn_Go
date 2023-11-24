package main

import (
	"leetcode_master"
)

func main() {
	head := &leetcode_master.ListNode{
		Val: 1,
	}
	head.Next = &leetcode_master.ListNode{
		Val: 2,
	}
	head.Next.Next = &leetcode_master.ListNode{
		Val: 3,
	}
	leetcode_master.ReverseList(head)
}
