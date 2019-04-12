package main

import (
	"fmt"
)

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(60, 7))
}

// 在上面程序中，第 7 行的 simple 函数返回了一个函数，并接受两个 int 参数，返回一个 int。
// 在第 15 行，我们调用了 simple 函数。
// 我们把 simple 的返回值赋值给了 s。现在 s 包含了 simple 函数返回的函数。
// 我们调用了 s，并向它传递了两个 int 参数（第 16 行）。
// 该程序输出 67。
