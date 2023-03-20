package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // 起始是 0，最後一位不取，結果：[3 5 7]

	var s1 []int = primes[3:]
	fmt.Println(s1) // 取第三位到最後，結果：[7 11 13]

	var s2 []int = primes[:2]
	fmt.Println(s2) // 取 0~2，結果：[2 3]
}
