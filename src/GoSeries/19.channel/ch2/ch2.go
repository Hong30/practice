package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}

// 在上述程序里，我们在第 12 行创建了一个 bool 类型的信道 done，并把 done 作为参数传递给了 hello 协程。
// 在第 14 行，我们通过信道 done 接收数据。
// 这一行代码发生了阻塞，除非有协程向 done 写入数据，否则程序不会跳到下一行代码。于是，
// 这就不需要用以前的 time.Sleep 来阻止 Go 主协程退出了。
// <-done 这行代码通过协程（译注：原文笔误，信道）done 接收数据，但并没有使用数据或者把数据存储到变量中。这完全是合法的。
// 现在我们的 Go 主协程发生了阻塞，等待信道 done 发送的数据。
// 该信道作为参数传递给了协程 hello，hello 打印出 Hello world goroutine，接下来向 done 写入数据。
// 当完成写入时，Go 主协程会通过信道 done 接收数据，于是它解除阻塞状态，打印出文本 main function。
