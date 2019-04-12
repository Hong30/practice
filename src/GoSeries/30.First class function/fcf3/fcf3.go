package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

// 在上面程序中，函数 appendStr 返回了一个闭包。
// 这个闭包绑定了变量 t。我们来理解这是什么意思。
// 在第 17 行和第 18 行声明的变量 a 和 b 都是闭包，它们绑定了各自的 t 值。
// 我们首先用参数 World 调用了 a。现在 a 中 t 值变为了 Hello World。
// 在第 20 行，我们又用参数 Everyone 调用了 b。
// 由于 b 绑定了自己的变量 t，因此 b 中的 t 还是等于初始值 Hello。
// 于是该函数调用之后，b 中的 t 变为了 Hello Everyone。程序的其他部分很简单，不再解释。

// 该程序会输出：

// Hello World
// Hello Everyone
// Hello World Gopher
// Hello Everyone !
