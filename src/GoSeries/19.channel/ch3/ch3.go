package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}

// 在上面程序里，我们向 hello 函数里添加了 4 秒的休眠（第 10 行）。
// 程序首先会打印 Main going to call hello go goroutine。
// 接着会开启 hello 协程，打印 hello go routine is going to sleep。
// 打印完之后，hello 协程会休眠 4 秒钟，而在这期间，主协程会在 <-done 这一行发生阻塞，等待来自信道 done 的数据。
// 4 秒钟之后，打印 hello go routine awake and going to write to done，接着再打印 Main received data。
