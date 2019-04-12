package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

// 在上述程序中，我们创建了容量为 1 的缓冲信道，并在第 18 行将它传入 increment 协程。
// 该缓冲信道用于保证只有一个协程访问增加 x 的临界区。
// 具体的实现方法是在 x 增加之前（第 8 行），传入 true 给缓冲信道。
// 由于缓冲信道的容量为 1，所以任何其他协程试图写入该信道时，都会发生阻塞，直到 x 增加后，信道的值才会被读取（第 10 行）。
// 实际上这就保证了只允许一个协程访问临界区。
// 该程序也输出：
// final value of x 1000
