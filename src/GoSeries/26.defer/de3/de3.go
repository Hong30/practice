package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

}

// 第 11 行，a 的初始值为 5。
// 在第 12 行执行 defer 语句的时候，由于 a 等于 5，因此延迟函数 printA 的实参也等于 5。
// 接着我们在第 13 行将 a 的值修改为 10。下一行会打印出 a 的值。
// 该程序输出：
// value of a before deferred function call 10
// value of a in deferred function 5
