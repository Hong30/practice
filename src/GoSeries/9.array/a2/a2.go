package main

import (
	"fmt"
)

func main() {
	var a [3]int //int array with length 3
	a[0] = 12    // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

// a[0] 将值赋给数组的第一个元素。

// package main

// import (
//     "fmt"
// )

// func main() {
//     a := [3]int{12, 78, 50} // short hand declaration to create array
//     fmt.Println(a)
// }
// 简短申明

// package main

// import (
//     "fmt"
// )

// func main() {
//     a := [3]int{12}
//     fmt.Println(a)
// }

// a := [3]int{12} 声明一个长度为 3 的数组，但只提供了一个值 12，剩下的 2 个元素自动赋值为 0。
// 这个程序将输出 [12 0 0]。
