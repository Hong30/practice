package main

import (
	"fmt"
)

func main() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

// 第 11 行，for range 循环会遍历一个字符串，并在第 12 行调用了 defer fmt.Printf("%c", v)。
// 这些延迟调用会添加到一个栈中，按照后进先出的顺序执行，因此，该字符串会逆序打印出来。
// 该程序会输出：
// Orignal String: Naveen
// Reversed String: neevaN
