package main

import "fmt"

func main() {
	test1()
	test2()
}

func test1() {
	fmt.Println("before panic")
	panic("crash")
	//unexe
	fmt.Println("before panic")

}

// 数组越界
func test2() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
}

// defer
func test3() {
	defer func() {
		fmt.Println("defer func")
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
}

// recover
func test4() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}
