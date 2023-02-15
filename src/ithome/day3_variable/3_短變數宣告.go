package main

import "fmt"

/*
 * 短變數宣告：
 * 在函數中，「:=」 簡潔賦值語句在明確類型的地方，可以替代 var 定義。
 *（「:=」 結構不能使用在函數外，函數外的每個語法都必須以關鍵字開始。）
 */
func main() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}