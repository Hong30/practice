package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
}

// 在第 11 行，go hello() 启动了一个新的 Go 协程。现在 hello() 函数与 main() 函数会并发地执行。
// 主函数会运行在一个特有的 Go 协程上，它称为 Go 主协程（Main Goroutine）。
// 该程序只会输出文本 main function。我们启动的 Go 协程究竟出现了什么问题？要理解这一切，我们需要理解两个 Go 协程的主要性质。
// 启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。
// 在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
// 如果希望运行其他 Go 协程，Go 主协程必须继续运行着。
// 如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
// 现在你应该能够理解，为何我们的 Go 协程没有运行了吧。
// 在第 11 行调用了 go hello() 之后，程序控制没有等待 hello 协程结束，立即返回到了代码下一行，打印 main function。
// 接着由于没有其他可执行的代码，Go 主协程终止，于是 hello 协程就没有机会运行了。
