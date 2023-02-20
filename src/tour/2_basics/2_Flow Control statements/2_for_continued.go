package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; { // 可以不須有初始值與後置語句
	// for sum < 1000 { // 與前一句同樣意思
		sum += sum
	}
	fmt.Println(sum)
}
