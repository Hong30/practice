package main

import (
	"fmt"
)

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}

// 程序的第 7 行，mutate 函数接收一个 rune 切片参数，它将切片的第一个元素修改为 'a'，
// 然后将 rune 切片转化为字符串，并返回该字符串。
// 程序的第 13 行调用了该函数。
// 我们把 h 转化为一个 rune 切片，并传递给了 mutate。这个程序输出 aello。
