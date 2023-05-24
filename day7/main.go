package main

import (
	"fmt"
	"gee"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Shon\n")
	})
	//panic
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"Shon"}
		c.String(http.StatusOK, names[100])
	})

	//end
	fmt.Println("Listening:9999")
	r.Run(":9999")
}
