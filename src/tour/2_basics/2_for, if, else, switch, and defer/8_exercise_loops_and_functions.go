package main

import (
	"fmt"
	"math"
) 

func sqrt(x float64) float64{
	z := 1.0
	for i := 1; i<10 ;i++{
		z -= (z*z - x) / (2*z)
		// fmt.Println("z = ", z)
	}
	return z
}

func main(){
	fmt.Println(math.Sqrt(4), sqrt(4))
	fmt.Println(math.Sqrt(10), sqrt(10))
}