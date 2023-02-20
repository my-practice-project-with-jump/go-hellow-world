package main

import "fmt"

// return 結果在後面
func add (x int, y int) int {
	return x+y
}

func main (){
	fmt.Print(add(42, 13))
}