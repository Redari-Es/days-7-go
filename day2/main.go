package main

import (
	"days-7-go/day2/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()
	// query
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s \n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	//end
	fmt.Println("Listening:9999")
	r.Run(":9999")
}
