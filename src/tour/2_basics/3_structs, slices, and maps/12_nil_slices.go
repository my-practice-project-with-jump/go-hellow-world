package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil { // 全空空，就是 nil
		fmt.Println("nil!")
	}
}
