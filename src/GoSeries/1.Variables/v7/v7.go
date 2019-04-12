package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)
}

// 变量也可以在运行时进行赋值。
