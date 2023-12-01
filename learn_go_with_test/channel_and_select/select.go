package main

import "fmt"

func Select1() {
	case1 := make(chan int)
	case2 := make(chan int)
	close(case1)
	close(case2)

	select {
	case <-case1:
		fmt.Print("case1")
	case case2 <- 1:
		fmt.Print("case2")
	default:
		fmt.Print("Default")
	}
}

func Select2() {
	c := make(chan int)

	done := false
	for !done {
		select {
		case <-c:
			fmt.Print("1")
			c = nil
		case c <- 1:
			fmt.Print("2")
		default:
			fmt.Print("3")
			done = true
		}
	}
}
