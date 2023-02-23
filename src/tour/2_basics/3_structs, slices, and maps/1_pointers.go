package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i               // 指向i
	fmt.Println(i, p, *p) // *p 通過指針讀取 i的值
	*p = 21               // 通過指針設置 i的值
	fmt.Println(i, p, *p) // *p 通過指針讀取 i的值

	p = &j // 指向j
	fmt.Println(j, p, *p)
	*p = *p / 37 // 通過指針對j進行除法運算
	fmt.Println(j, p, *p)

}
