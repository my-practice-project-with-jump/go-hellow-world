package main

/*
 * 等同：
 * import "fmt"
 * improt "math"
 */
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\nSo, what would like to do?", math.Sqrt(7)) // 開根號
}
