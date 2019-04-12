package main

import (
	"fmt"
)

func main() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
}

// 可以忽略声明数组的长度，并用 ... 代替，让编译器为你自动计算长度
