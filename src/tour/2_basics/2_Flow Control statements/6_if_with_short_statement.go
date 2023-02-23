package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // 因與下一行的 ')' 在不同行，所以會被認為之後可能還有資料，若後面直接加 ) 則不需要,，可以參考下一個範例
	)
}
