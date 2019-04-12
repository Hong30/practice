package main

import (
	"fmt"
)

func main() {
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
}

// 切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。
// 可以使用 append 函数将值追加到 nil 切片。
// 在上面的程序 names 是 nil，我们已经添加 3 个字符串给 names。
