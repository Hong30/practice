package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// 在程序的第 9 行，我们试图打开路径为 /test.txt 的文件（playground 显然并不存在这个文件）。os 包里的 Open 函数有如下签名：
// func Open(name string) (file *File, err error)

// 如果成功打开文件，Open 函数会返回一个文件句柄（File Handler）和一个值为 nil 的错误。
// 而如果打开文件时发生了错误，会返回一个不等于 nil 的错误。
// 如果一个函数 或方法 返回了错误，按照惯例，错误会作为最后一个值返回。
// 于是 Open 函数也是将 err 作为最后一个返回值。
// 按照 Go 的惯例，在处理错误时，通常都是将返回的错误与 nil 比较。
// nil 值表示了没有错误发生，而非 nil 值表示出现了错误。
// 在这里，我们第 10 行检查了错误值是否为 nil。
// 如果不是 nil，我们会简单地打印出错误，并在 main 函数中返回。

// 运行该程序会输出：

// open /test.txt: No such file or directory
