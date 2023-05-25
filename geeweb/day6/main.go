package main

import (
	"fmt"
	"gee"
	"html/template"
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
	r := gee.New()
	r.Use(gee.Logger()) //global middleware
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	stu1 := &student{Name: "RedAries小行星", Age: 20}
	stu2 := &student{Name: "Geektutu", Age: 20}
	// stu2 := &student{Name: "Geektutu", Age: 20}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	//get
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2023, 4, 11, 0, 0, 0, 0, time.UTC),
		})
	})
	//panic
	// r.GET("/panic", func(c *gee.Context) {
	// 	names := []string{"shon"}
	// 	c.String(http.StatusOK, names[100])
	// })

	//end
	fmt.Println("Listening:9999")
	r.Run(":9999")
}
