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

// 单词搜索Ⅱ
type Dic struct {
	children map[byte]*Dic
	word     string
}

func (dic *Dic) Insert(word string) {
	node := dic
	for i := range word {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = &Dic{
				children: map[byte]*Dic{},
			}
		}
		node = node.children[ch]
	}
	node.word = word
}

func FindWords(board [][]byte, words []string) []string {
	dic := &Dic{children: map[byte]*Dic{}}
	m, n := len(board), len(board[0])
	result := []string{}
	for _, word := range words {
		dic.Insert(word)
	}

	var dfs func(i int, j int, node *Dic)
	dfs = func(i, j int, node *Dic) {
		ch := board[i][j]
		nxt := node.children[ch]
		if nxt == nil {
			// 不含当前词
			return
		}
		if nxt.word != "" {
			// 一个词遍历完则加入到结果
			result = append(result, nxt.word)
			// 删除该词避免重复取到
			nxt.word = ""
		}

		if len(nxt.children) > 0 {
			// 非根节点
			next := [][]int{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}}
			board[i][j] = '#'
			for _, nextPoint := range next {
				x, y := nextPoint[0], nextPoint[1]
				if x < 0 || x >= m || y < 0 || y >= n || board[x][y] == '#' {
					// 触碰边界则跳过
					continue
				}
				dfs(x, y, nxt)
			}
			board[i][j] = ch
		}

		if len(nxt.children) == 0 {
			// 如果某一个子路径上所有单词都被读取，则进行剪枝操作
			delete(node.children, ch)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, dic)
		}
	}
	return result
}
