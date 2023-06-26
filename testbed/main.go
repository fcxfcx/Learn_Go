package main

import (
	"leetcode"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	result := leetcode.EvalRPN(tokens)
	println(result)
}
