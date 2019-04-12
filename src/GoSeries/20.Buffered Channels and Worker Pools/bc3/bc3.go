package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 在上面程序里，我们向容量为 2 的缓冲信道写入 3 个字符串。
// 当在程序控制到达第 3 次写入时（第 11 行），由于它超出了信道的容量，因此这次写入发生了阻塞。
// 现在想要这次写操作能够进行下去，必须要有其它协程来读取这个信道的数据。
// 但在本例中，并没有并发协程来读取这个信道，因此这里会发生死锁（deadlock）。
// 程序会在运行时触发 panic，信息如下：
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//     /tmp/sandbox274756028/main.go:11 +0x100
