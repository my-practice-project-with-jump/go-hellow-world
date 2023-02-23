package main

import "fmt"

// 參考 https://blog.go-zh.org/defer-panic-and-recover
func main() {
	fmt.Println("counting")

	// 按照後進先出
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
