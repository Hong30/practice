package main

import (
	"fmt"
)

func main() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

// 程序中，b 初始化为 nil，接着将 a 的地址赋值给 b。
