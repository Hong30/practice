package main

import (
	"fmt"
)

func main() {
	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v", a, a)
}

// 5.9 在语法中是浮点型，8 是整型，5.9/8 是允许的，因为两个都是数字常量。
// 除法的结果是 0.7375 是一个浮点型，所以 a 的类型是浮点型。
