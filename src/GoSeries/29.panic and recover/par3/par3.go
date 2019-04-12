package main

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}

// 在上面的程序中，函数 b() 在第 23 行发生 panic。
// 函数 a() 调用了一个延迟函数 recovery()，用于恢复 panic。在第 17 行，函数 b() 作为一个不同的协程来调用。
// 下一行的 Sleep 只是保证 a() 在 b() 运行结束之后才退出。

// 你认为程序会输出什么？panic 能够恢复吗？答案是否定的，panic 并不会恢复。
// 因为调用 recovery 的协程和 b() 中发生 panic 的协程并不相同，因此不可能恢复 panic。

// 运行该程序会输出：

// Inside A
// Inside B
// panic: oh! B panicked

// goroutine 5 [running]:
// main.b()
//     /tmp/sandbox388039916/main.go:23 +0x80
// created by main.a
//     /tmp/sandbox388039916/main.go:17 +0xc0
// 从输出可以看出，panic 没有恢复。

// 如果函数 b() 在相同的协程里调用，panic 就可以恢复。

// 如果程序的第 17 行由 go b() 修改为 b()，就可以恢复 panic 了，因为 panic 发生在与 recover 相同的协程里。
// 如果运行这个修改后的程序，会输出：

// Inside A
// Inside B
// recovered: oh! B panicked
// normally returned from main
