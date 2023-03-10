package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // 等同 switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
