package main

import (
	"fmt"
	"math"
)

func main() {
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Println(i, f, u)

	var x, y int = 2, 4
	var fl float64 = math.Sqrt((float64(x*x + y*y)))
	var z uint  = uint(f)
	fmt.Println(x, y, fl, z)
}