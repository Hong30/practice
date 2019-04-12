package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}

// 程序的第 7 行，describe(i interface{}) 函数接收空接口作为参数，因此，可以给这个函数传递任何类型。
// 在第 13 行、第 15 行和第 21 行，我们分别给 describe 函数传递了 string、int 和 struct。
