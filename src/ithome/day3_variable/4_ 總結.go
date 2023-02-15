package main

import "fmt"

// var a // 不定型別的變數， 會出錯，不能不宣告型別
var b int // 宣告成 int
var c int = 10 // 初始化同時宣告
var d, e int // d 跟 e 都是 int
// var f int, g string // 會出錯，不同型別不能在同一行

var (
	h bool = false // 記得要不同行不然會錯
	i int
	j = "hello"
)

func main() {
	k := 0
	l, m, n := 0, true, "tacolin" // 這樣就可以不同型別寫在同一行

	// fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l. m, n)
	fmt.Println( b, c, d, e, h, i, j, k, l, m, n)
}