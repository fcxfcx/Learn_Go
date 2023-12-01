package main

import (
	"fmt"
	"time"
)

func print(u <-chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("print int", <-u)
}

func Channel1() {
	c := make(chan int, 5)
	a := 0

	c <- a
	fmt.Println(a)
	// modify a
	a = 1

	go print(c)
	time.Sleep(5 * time.Second)
	fmt.Println(a)
}

type people struct {
	name string
}

var u = people{name: "A"}

func printPeople(u <-chan *people) {
	time.Sleep(2 * time.Second)
	a := <-u
	fmt.Printf("printPeople, %p\n", a)
}

func Channel2() {
	c := make(chan *people, 5)
	var a = &u
	c <- a
	fmt.Printf("%p\n", a)
	// modify a
	a = &people{name: "B"}

	go printPeople(c)
	time.Sleep(5 * time.Second)
	fmt.Printf("%p", a)
}
