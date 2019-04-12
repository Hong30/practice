package main

import (
	"fmt"
)

func main() {
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

// if num％2 == 0 语句检测 num 取 2 的余数是否为零。
// 如果是为零则打印输出 "the number is even"，
// 如果不为零则打印输出 "the number is odd"。
