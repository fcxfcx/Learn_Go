package main

import (
	"leetcode"
)

func main() {
	nums1 := [12]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := leetcode.Trap(nums1[:])
	println(result)
}
