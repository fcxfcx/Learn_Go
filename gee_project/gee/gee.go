package gee

import (
	"fmt"
	"log"
	"net/http"
)

// 处理函数类
type HandlerFunc func(c *Context)

// Gee引擎
type Engine struct {
	*RouterGroup // 使用嵌套类型，因为engine本身是一个最大的路由组
	router       *router
	groups       []*RouterGroup
}

// 初始化构建方法
func New() *Engine {
	engine := &Engine{router: newRouter()}
	// engine继承了RouterGroup，它的engine属性指向它自己
	// 这样在使用过程中可以在engine中添加路由也可以在某一个分组下添加路由了
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 启动方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 寻找对应的处理器并对请求进行回应
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	fmt.Println(c)
	engine.router.handle(c)
}

// 分组路由
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine // 给分组访问engine的方式，方便间接访问各个接口
}

// 创建新的分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		engine: engine,
		prefix: group.prefix + prefix,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 在路由组下添加新的路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {

	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// 添加GET请求下的路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// 添加GET请求下的路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
