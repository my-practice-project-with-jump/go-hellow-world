package main

import (
	"fmt"
	"time"
)

/*
碰到符合的條件，該區塊執行結束即跳脫switch區域
*/
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
