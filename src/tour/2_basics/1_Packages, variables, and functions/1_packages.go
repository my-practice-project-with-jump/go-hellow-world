package main

import (
	"fmt"
	"math/rand"
)

// 關於 math/rand 說明，請參考：https://pkg.go.dev/math/rand#Intn
func main() {
	// 回傳一個隨機的1~10之間的正整數
	fmt.Print("My favorite number is ", rand.Intn(10)) 
}