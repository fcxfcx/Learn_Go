package gee

import "strings"

// 前缀树路由
type node struct {
	pattern  string  // 匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	// 插入操作，使用回溯
	// pattern是待匹配的路径，parts是路径拆分成每一部分
	// height是当前正在匹配的部分的索引（也是路径深度）
	if len(parts) == height {
		// 匹配完成后，当前节点对应的pattern就确定了
		// 可以认为是一个停止信号，意味着这里的节点是一个pattern的结果
		n.pattern = pattern
		return
	}
	// 当前需要匹配的节点
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 如果还没有对应的子节点则新建一个
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 已有新子节点就继续按子节点路径走
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 查找操作，parts是需要查找的路径，height是当前的部分对应的索引
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	// 搜索所有满足条件的子路径
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
