package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

// 在上述程序里，第 7 行的 increment 函数把 x 的值加 1，并调用 WaitGroup 的 Done()，通知该函数已结束。
// 在上述程序的第 15 行，我们生成了 1000 个 increment 协程。
// 每个 Go 协程并发地运行，由于第 8 行试图增加 x 的值，因此多个并发的协程试图访问 x 的值，这时就会发生竞态条件。
// 由于 playground 具有确定性，竞态条件不会在 playground 发生，请在你的本地运行该程序。
// 请在你的本地机器上多运行几次，可以发现由于竞态条件，每一次输出都不同。
// 我其中遇到的几次输出有 final value of x 941、final value of x 928、final value of x 922 等。
