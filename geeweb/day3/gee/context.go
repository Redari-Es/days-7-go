package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 设置别名用于gee.H
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// 提供路由参数的访问
	Params map[string]string
	// response info
	StatusCode int
}

// newContext
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 提供访问Query和PostForm参数的方法
// PostForm
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// query
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

// param
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
