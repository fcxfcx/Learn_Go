package main

import "fmt"

func main() {
	// tips: select test1可能会panic
	fmt.Println("select test 1:")
	Select1()
	fmt.Println("\n---------------")

	fmt.Println("select test 2:")
	Select2()
	fmt.Println("\n---------------")

	fmt.Println("channel test 1:")
	Channel1()
	fmt.Println("\n---------------")

	fmt.Println("channel test 2:")
	Channel2()
	fmt.Println("\n---------------")
}
