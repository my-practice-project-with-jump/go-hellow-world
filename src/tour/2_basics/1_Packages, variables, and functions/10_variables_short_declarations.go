package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // := 此種寫法只能在function內，function 外一定要var開頭
	c, python, java := true, false, "no!" // 會推論出它的型態

	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", k, k)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", python, python)
	fmt.Printf("Type: %T Value: %v\n", java, java)
}