package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // 擷取切片使其長度為0
	printSlice(s)

	s = s[:4] // 拓展其長度
	printSlice(s)

	s = s[2:] // 捨棄前兩值
	printSlice(s) // 往前擷取，資料會不見

	s = s[:] // 捨棄前兩值
	printSlice(s) 

	s = s[:4] // 捨棄前兩值
	printSlice(s) // 往後擷取資料會保留


}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}