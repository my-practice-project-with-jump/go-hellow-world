package main

import "fmt"

// 定義變數時可以直接賦予初始值，變數與賦予的值要互相對應。
var x, y, z int = 1, 2, 3

// 如果有初始化的話，型別就可以省略；變數會直接取用初始化的類型
var c, python, java = true, false, "no!"

func main (){
	fmt.Println(x, y, z, c, python, java)
}