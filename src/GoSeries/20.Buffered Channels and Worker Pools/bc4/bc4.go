package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

// 在上面的程序里，我们创建了一个容量为 3 的信道，于是它可以保存 3 个字符串。接下来，
// 我们分别在第 9 行和第 10 行向信道写入了两个字符串。
// 于是信道有两个字符串排队，因此其长度为 2。在第 13 行，我们又从信道读取了一个字符串。
// 现在该信道内只有一个字符串，因此其长度变为 1。该程序会输出：
// capacity is 3
// length is 2
// read value naveen
// new length is 1
