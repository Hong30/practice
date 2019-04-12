// package main
// import (
//     "fmt"
// )

// func main() {
//     b := 255
//     a := &b
//     fmt.Println("address of b is", a)
//     fmt.Println("value of b is", *a)
// }

// 程序的第 10 行，我们将 a 解引用，并打印了它的值。
// 不出所料，我们会打印出 b 的值。

package main

import (
	"fmt"
)

func main() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
}

// 把 a 指向的值加 1，由于 a 指向了 b，因此 b 的值也发生了同样的改变。
// 于是 b 的值变为 256。
