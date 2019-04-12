package main

import (
	"fmt"
)

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

// 在上面的实例中，第 7 行我们定义了一个函数 simple，simple 接收一个函数参数（该函数接收两个 int 参数，返回一个 a 整型）。
// 在 main 函数的第 12 行，我们创建了一个匿名函数 f，其签名符合 simple 函数的参数。
// 我们在下一行调用了 simple，并传递了参数 f。该程序打印输出 67。
