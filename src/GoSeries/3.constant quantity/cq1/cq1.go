package main

func main() {
	const a = 55 // 允许
	// a = 89       // 不允许重新赋值
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println("Hello, playground")
// 	var a = math.Sqrt(4)   // 允许
// 	const b = math.Sqrt(4) // 不允许
// }

// 常量的值会在编译的时候确定。
// 因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。

// a 是变量，因此我们可以将函数 math.Sqrt(4) 的返回值赋值给它。
// b 是一个常量，它的值需要在编译的时候就确定。
// 函数 math.Sqrt(4) 只会在运行的时候计算，因此 const b = math.Sqrt(4) 会出错
