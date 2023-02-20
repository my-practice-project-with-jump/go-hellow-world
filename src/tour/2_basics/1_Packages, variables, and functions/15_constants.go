package main

/*
constant(常數)
const 開頭
:=  不可為 constant
*/

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happly", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Printf("Pi type: %T\nWorld type: %T\nTruth type: %T\n", Pi, World, Truth)
}