package main

import (
	"fmt"
	"math"
)

/*
 * math.Sqrt: 開根號
 * fmt.Sprint: 組合資料內容
 */

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
