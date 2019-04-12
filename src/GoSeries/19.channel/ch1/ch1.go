package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}

// 由于信道的零值为 nil，在第 6 行，信道 a 的值就是 nil。
// 于是，程序执行了 if 语句内的语句，定义了信道 a。
// 程序中 a 是一个 int 类型的信道。
