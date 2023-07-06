package main

import (
	"leetcode"
)

func main() {
	cache := leetcode.LRUConstructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Get(2)
	cache.Put(1, 1)
	cache.Put(4, 1)
	cache.Get(2)
}
