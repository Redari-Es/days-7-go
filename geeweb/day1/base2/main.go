package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := new(Engine)
	fmt.Println("Listenning: 9999")
	log.Fatal(http.ListenAndServe(":9999", engine))
}

// Engine is the uni handler for all request
// 该引擎结构用于处理所有访问
type Engine struct{} //空结构体

// 实现方法
// ResponWriter可以构造针对该请求的响应
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 路径状态机，拦截http请求，统一处理
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
