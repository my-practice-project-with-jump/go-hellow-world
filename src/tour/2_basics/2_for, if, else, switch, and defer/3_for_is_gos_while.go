package main

import "fmt"

func main() {
	sum := 1
	for sum < 60 {
		sum += sum
	}
	fmt.Println(sum)

	// 與上述同樣意思
	sum2 := 1
	for ; sum2 < 60; sum2 += sum2 {
	}
	fmt.Println(sum2)
}
