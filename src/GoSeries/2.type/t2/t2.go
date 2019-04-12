package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
}

// 推断出 a 和 b 为 int 类型，且大小都是 32 位（4 字节）。
// 如果你在 64 位系统上运行上面的代码，会有不同的输出。在 64 位系统下，a 和 b 会占用 64 位（8 字节）的大小。
