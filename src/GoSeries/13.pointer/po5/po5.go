// package main

// import (
//     "fmt"
// )

// func modify(arr *[3]int) {
//     (*arr)[0] = 90
// }

// func main() {
//     a := [3]int{89, 90, 91}
//     modify(&a)
//     fmt.Println(a)
// }

// 上面程序的第 13 行中，我们将数组的地址传递给了 modify 函数。
// 在第 8 行，我们在 modify 函数里把 arr 解引用，并将 90 赋值给这个数组的第一个元素。
// 程序会输出 [90 90 91]。

// a[x] 是 (*a)[x] 的简写形式，因此上面代码中的 (*arr)[0] 可以替换为 arr[0]。
// 下面我们用简写形式重写以上代码。
// package main

// import (
//     "fmt"
// )

// func modify(arr *[3]int) {
//     arr[0] = 90
// }

// func main() {
//     a := [3]int{89, 90, 91}
//     modify(&a)
//     fmt.Println(a)
// }

// 接下来我们用切片来重写之前的代码。
package main

import (
	"fmt"
)

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	a := [3]int{89, 90, 91}
	modify(a[:])
	fmt.Println(a)
}

// 我们将一个切片传递给了 modify 函数。
// 在 modify 函数中，我们把切片的第一个元素修改为 90。程序也会输出 [90 90 91]。
// 所以别再传递数组指针了，而是使用切片吧。上面的代码更加简洁，也更符合 Go 语言的习惯。
