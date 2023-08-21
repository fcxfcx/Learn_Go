package main

import (
	"leetcode"
)

func main() {
	trie := leetcode.WDConstructor()
	trie.AddWord("a")
	trie.AddWord("a")
	trie.Search(".")
	trie.Search("a")
	trie.Search("aa")
	trie.Search(".a")
	trie.Search("a.")
}
