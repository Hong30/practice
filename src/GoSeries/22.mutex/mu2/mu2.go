package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

// Mutex 是一个结构体类型，我们在第 15 行创建了 Mutex 类型的变量 m，其值为零值。
// 在上述程序里，我们修改了 increment 函数，将增加 x 的代码（x = x + 1）放置在 m.Lock() 和 m.Unlock()之间。
// 现在这段代码不存在竞态条件了，因为任何时刻都只允许一个协程执行这段代码。
// 于是如果运行该程序，会输出：
// final value of x 1000
// 在第 18 行，传递 Mutex 的地址很重要。如果传递的是 Mutex 的值，而非地址，那么每个协程都会得到 Mutex 的一份拷贝，竞态条件还是会发生。
