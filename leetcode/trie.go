package leetcode

// 前缀树
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func TrieConstructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	node := trie
	for _, b := range word {
		b -= 'a'
		if node.children[b] == nil {
			node.children[b] = &Trie{}
		}
		node = node.children[b]
	}
	node.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	node := trie.searchPrefix(word)
	return node != nil && node.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	return trie.searchPrefix(prefix) != nil
}

func (trie *Trie) searchPrefix(prefix string) *Trie {
	node := trie
	for _, b := range prefix {
		b -= 'a'
		if node.children[b] == nil {
			return nil
		}
		node = node.children[b]
	}
	return node
}

// 添加与搜索单词
type WordDictionary struct {
	children map[byte]*WordDictionary
	isEnd    bool
}

func WDConstructor() WordDictionary {
	return WordDictionary{
		children: map[byte]*WordDictionary{},
	}
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd
	for _, b := range []byte(word) {
		if _, ok := node.children[b]; !ok {
			node.children[b] = &WordDictionary{
				children: map[byte]*WordDictionary{},
			}
		}
		node = node.children[b]
	}
	node.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	var dfs func(word string, node *WordDictionary) bool
	dfs = func(word string, node *WordDictionary) bool {
		if len(word) == 0 {
			return node.isEnd
		}
		char := word[0]
		if char == '.' {
			for _, child := range node.children {
				if dfs(word[1:], child) {
					return true
				}
			}
		} else {
			if _, ok := node.children[char]; ok {
				return dfs(word[1:], node.children[char])
			}
		}
		return false
	}
	return dfs(word, wd)
}
