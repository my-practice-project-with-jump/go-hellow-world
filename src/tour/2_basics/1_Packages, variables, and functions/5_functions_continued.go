package main

import "fmt"

// 連續兩個相同型別的參數，前面可以忽略不寫
func add (x, y int) int {
	return x + y
}

func main () {
	fmt.Print(add(42, 13))
}