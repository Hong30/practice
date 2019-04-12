package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

// 在上面程序的第 13 行，我们调用了 time 包里的函数 Sleep，该函数会休眠执行它的 Go 协程。
// 在这里，我们使 Go 主协程休眠了 1 秒。
// 因此在主协程终止之前，调用 go hello() 就有足够的时间来执行了。
// 该程序首先打印 Hello world goroutine，等待 1 秒钟之后，接着打印 main function。
