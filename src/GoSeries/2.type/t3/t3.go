package main

import (
	"fmt"
)

func main() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

// a 和 b 的类型为 float64（float64 是浮点数的默认类型）。
// 我们把 a 和 b 的和赋值给变量 sum，把 b 和 a 的差赋值给 diff，接下来打印 sum 和 diff。
// no1 和 no2 也进行了相同的计算。
