package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}

// 上述程序中，第 8 行的 process 函数休眠了 10500 毫秒（10.5 秒），接着把 process successful 写入 ch 信道。
// 在程序中的第 15 行，并发地调用了这个函数。
// 在并发地调用了 process 协程之后，主协程启动了一个无限循环。
// 这个无限循环在每一次迭代开始时，都会先休眠 1000 毫秒（1 秒），然后执行一个 select 操作。
// 在最开始的 10500 毫秒中，由于 process 协程在 10500 毫秒后才会向 ch 信道写入数据，
// 因此 select 语句的第一个 case（即 case v := <-ch:）并未就绪。
// 所以在这期间，程序会执行默认情况，该程序会打印 10 次 no value received。
// 在 10.5 秒之后，process 协程会在第 10 行向 ch 写入 process successful。
// 现在，就可以执行 select 语句的第一个 case 了，
// 程序会打印
// received value: process successful，
// 然后程序终止。
