package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // 中間不會有 ()
		sum += i
	}
	fmt.Println(sum)
}
