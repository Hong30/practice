package main

import (
	"fmt"
)

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}

// & 操作符用于获取变量的地址。
// 上面程序的第 9 行我们把 b 的地址赋值给 *int 类型的 a。
// 我们称 a 指向了 b。当我们打印 a 的值时，会打印出 b 的地址。
