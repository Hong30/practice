package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)
}

// 语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片。
// 因此，在上述程序的第 9 行中, a[1:4] 从索引 1 到 3 创建了 a 数组的一个切片表示。因此, 切片 b 的值为 [77 78 79]。

// package main

// import (
//     "fmt"
// )

// func main() {
//     c := []int{6, 7, 8} // creates and array and returns a slice reference
//     fmt.Println(c)
// }

// 另一种创建切片的方法
// c：= [] int {6，7，8} 创建一个有 3 个整型元素的数组，并返回一个存储在 c 中的切片引用。
