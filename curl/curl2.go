package main

import (
	"fmt"
	"os/exec"
)

func main() {
	urls := []string{
		"https://www.example.com/path1",
		"https://www.example.com/path2",
		"https://www.example.com/path3",
	}
	method := "GET" // or POST, etc.

	for _, url := range urls {
		curlCmd := exec.Command("curl", method, url)
		err := curlCmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Curl command executed successfully")
	}
}
