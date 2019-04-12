// package main

// import (
// 	"fmt"
// )

// func find(num int, nums ...int) {
// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }
// func main() {
// 	nums := []int{89, 90, 95}
// 	find(89, nums)
// }

// 在第 23 行中，我们将一个切片传给一个可变参数函数。
// 这种情况下无法通过编译，编译器报出错误 main.go:23: cannot use nums (type []int) as type int in argument to find
// 由可变参数函数的定义可知，nums ...int 意味它可以接受 int 类型的可变参数。

// 在上面程序的第 23 行，nums 作为可变参数传入 find 函数。
// 前面我们知道，这些可变参数参数会被转换为 int 类型切片然后在传入 find 函数中。
// 但是在这里 nums 已经是一个 int 类型切片，编译器试图在 nums 基础上再创建一个切片，

// 这里之所以会失败是因为 nums 是一个 []int类型 而不是 int类型。
// 那么有没有办法给可变参数函数传入切片参数呢？答案是肯定的。
// 有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。
// 如果这样做，切片将直接传入函数，不再创建新的切片
// 在上面的程序中，如果你将第 23 行的 find(89, nums) 替换为 find(89, nums...) ，程序将成功编译

package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	nums := []int{89, 90, 95}
	find(89, nums...)
}
