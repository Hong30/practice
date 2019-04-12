package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// WaitGroup 是一个结构体类型，我们在第 18 行创建了 WaitGroup 类型的变量，其初始值为零值。WaitGroup 使用计数器来工作。
// 当我们调用 WaitGroup 的 Add 并传递一个 int 时，WaitGroup 的计数器会加上 Add 的传参。要减少计数器，
// 可以调用 WaitGroup 的 Done() 方法。Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。
// 上述程序里，for 循环迭代了 3 次，我们在循环内调用了 wg.Add(1)（第 20 行）。因此计数器变为 3。
// for 循环同样创建了 3 个 process 协程，然后在第 23 行调用了 wg.Wait()，确保 Go 主协程等待计数器变为 0。
// 在第 13 行，process 协程内调用了 wg.Done，可以让计数器递减。一旦 3 个子协程都执行完毕（即 wg.Done() 调用了 3 次），
// 那么计数器就变为 0，于是主协程会解除阻塞。
// 在第 21 行里，传递 wg 的地址是很重要的。
// 如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，main 函数并不会知道。
// 该程序输出：
// started Goroutine  2
// started Goroutine  0
// started Goroutine  1
// Goroutine 0 ended
// Goroutine 2 ended
// Goroutine 1 ended
// All go routines finished executing
