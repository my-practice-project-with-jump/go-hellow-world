package main

import "fmt"

// 回傳多個參數
func swap (x, y string) (string ,string) {
	return y, x
}

func main () {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}