package main

import "fmt"

func main() {
	// defer: 將函數推延到外成函數返回後執行。
	defer fmt.Println("world")

	fmt.Println("hello")
}
