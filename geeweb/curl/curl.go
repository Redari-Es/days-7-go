package main

import (
	"fmt"
	"os/exec"
)

func main() {
	curlURL := "http://localhost:9999"
	curlCmd := exec.Command("curl", curlURL)
	err := curlCmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Curl command executed successfully")
}
