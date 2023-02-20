package main

import "fmt"

func main() {
	v := 42 // int
	fmt.Printf("v is of Type %T value %v\n", v, v)
	f := 3.142        // float64
	fmt.Printf("f is of Type %T value %v\n", f, f)
	g := 0.867 + 0.5i // complex128
	fmt.Printf("g is of Type %T value %v\n", g, g)
}