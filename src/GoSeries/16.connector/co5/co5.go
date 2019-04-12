// package main

// import (
//     "fmt"
// )

// func assert(i interface{}) {
//     s := i.(int) //get the underlying int value from i
//     fmt.Println(s)
// }
// func main() {
//     var s interface{} = 56
//     assert(s)
// }

// 在第 12 行，s 的具体类型是 int。在第 8 行，我们使用了语法 i.(int) 来提取 i 的底层 int 值。
// 该程序会打印 56。

// 在上面程序中，如果具体类型不是 int，会发生什么呢？

// package main

// import (
//     "fmt"
// )

// func assert(i interface{}) {
//     s := i.(int)
//     fmt.Println(s)
// }
// func main() {
//     var s interface{} = "Steven Paul"
//     assert(s)
// }

// 在上面程序中，我们把具体类型为 string 的 s 传递给了 assert 函数，试图从它提取出 int 值。
// 该程序会报错：panic: interface conversion: interface {} is string, not int.。
// 要解决该问题，我们可以使用以下语法：
// v, ok := i.(T)
// 如果 i 的具体类型是 T，那么 v 赋值为 i 的底层值，而 ok 赋值为 true。
// 如果 i 的具体类型不是 T，那么 ok 赋值为 false，v 赋值为 T 类型的零值，此时程序不会报错。

package main

import (
	"fmt"
)

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}

// 当给 assert 函数传递 Steven Paul 时，由于 i 的具体类型不是 int，
// ok 赋值为 false，而 v 赋值为 0（int 的零值）。
