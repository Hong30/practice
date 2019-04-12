package main

import (
	"fmt"
)

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", nos) // modifications are visible outside
}

// 上述程序的行号 17 中，调用函数将切片中的每个元素递减 2。
// 在函数调用后打印切片时，这些更改是可见的。
// 如果你还记得，这是不同于数组的，对于函数中一个数组的变化在函数外是不可见的。
