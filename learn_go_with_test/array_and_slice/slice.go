package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := make([]int, 5)
	case1(s)
	case2(s)
	PrintSliceStruct(&s)
}

// 底层数组不变
func case1(s []int) {
	s[1] = 1
	PrintSliceStruct(&s)
}

// 底层数组变化
func case2(s []int) {
	s = append(s, 0)
	s[1] = 1
	PrintSliceStruct(&s)
}

func PrintSliceStruct(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("slice struct: %+v, slice is %v \n", ss, s)
}
