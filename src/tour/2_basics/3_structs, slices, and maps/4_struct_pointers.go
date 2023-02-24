package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // v 跟 p 是指相同一個東西，只是一個是值，一個是pointer
	p.X = 1e9
	fmt.Println(v, p)
	v.X = 1e3
	fmt.Println(v, p)
	q := v // q 跟 v 是不同東西
	q.Y = 5
	fmt.Println(q, v)

}
