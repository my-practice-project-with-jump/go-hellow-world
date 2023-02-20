package main

import "fmt"

const (
	// 將1左一100位來創建一個非常大的數字
	// 即這個數的二進制是1後面跟著100個0
	Big = 1 << 100
	// 再往右一99位，即Small = 1 << 1，或者說 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 +1 }

func needFlat(x float64) float64{
	return x * 0.1
}

func main() {
	fmt.Printf("Big type = %T, value = %v\n", Small, Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFlat(Small))
	fmt.Print(needFlat(Big))

	// 會噴這個錯誤: cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Printf (overflows)
	// 因該是因為值過大，故無法指定為 int
	// fmt.Printf("Big type = %T, value = %f\n", Big, Big)
	// fmt.Print(needInt(Big))
}
