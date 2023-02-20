package main 

import "fmt"

// return 被命名的參數，一般不建議這樣做，會影響閱讀性
func split (sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 没有参数的 return 语句返回已命名的返回值
}

func main () {
	fmt.Println(split(17))
}