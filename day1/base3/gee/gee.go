package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
// 提供框架给用户，用例定义路由映射的处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
// 添加一张路由映射表，key有请求方法和静态路由地址构成,value是用户映射的处理方法
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine 构造器
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// 当用户调用请求方法时，会将路由和处理方法注册到映射表router中
// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// Post defines the method to add POSTrequest
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
// ListenAndServe的包装
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 解析请求的路径，查找路由映射表，查到则执行注册的处理方法。查不到返回404
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
