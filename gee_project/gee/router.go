package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node       // 以前缀树形式存储的路由路径，key为方法GET/POST等
	handlers map[string]HandlerFunc // 对应的处理方法，key为整个请求例如'GET-/p/:lang/doc'
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	// 待匹配路径分段
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				// 遇到*则后面都可匹配，不用再讨论
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		// 如果路由里还没有对应的请求类型则需要新建
		r.roots[method] = &node{}
	}
	// 在对应的请求类型节点下添加pattern，构造前缀树
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	// 首先找到请求方法对应的根节点
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	// 从节点开始搜索path对应的节点
	n := root.search(searchParts, 0)

	if n != nil {
		// 这里的parts是可能含有模糊匹配成功的路径的
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			// 对于模糊匹配成功的部分，需要确定匹配的具体数据是什么
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				// 如果模糊匹配到了*则path后面所有的数据都是当前的param
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
