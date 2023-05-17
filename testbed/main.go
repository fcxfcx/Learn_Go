package main

import (
	"leetcode"
)

func main() {
	nums1 := [7]int{2, 2, 1, 1, 1, 2, 2}
	println(leetcode.MajorityElement(nums1[:]))
}
