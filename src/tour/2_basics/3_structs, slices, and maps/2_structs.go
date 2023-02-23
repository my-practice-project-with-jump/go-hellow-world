package main

import "fmt"

// A struct is a collection of fields.
// 類似 java 物件 ?
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
