package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:5]
	fmt.Println(s) // 3, 5, 7, 11

	s = s[:3] 
	fmt.Println(s) // 3, 5, 7

	s = s[1:]
	fmt.Println(s) // 5, 7

	s = s[:]
	fmt.Println(s) // 5, 7
}