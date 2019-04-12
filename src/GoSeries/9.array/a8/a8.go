package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}

// 第 17 行，用简略语法声明一个二维字符串数组 a 。
// 20 行末尾的逗号是必需的。
// 这是因为根据 Go 语言的规则自动插入分号。
// 另外一个二维数组 b 在 23 行声明，字符串通过每个索引一个一个添加。
// 这是另一种初始化二维数组的方法。
// 第 7 行的 printarray 函数使用两个 range 循环来打印二维数组的内容。
