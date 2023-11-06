package main

import (
	"top_100_liked"
)

func main() {
	lc := top_100_liked.Constructor(2)
	lc.Put(2, 1)
	lc.Put(2, 2)
	lc.Get(2)
	lc.Put(1, 1)
	lc.Put(4, 1)
	lc.Get(2)
}
