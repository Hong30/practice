// package main

// import "fmt"

// func sendData(sendch chan<- int) {
//     sendch <- 10
// }

// func main() {
//     sendch := make(chan<- int)
//     go sendData(sendch)
//     fmt.Println(<-sendch)
// }

// 上面程序的第 10 行，我们创建了唯送（Send Only）信道 sendch。chan<- int 定义了唯送信道，
// 因为箭头指向了 chan。在第 12 行，我们试图通过唯送信道接收数据，于是编译器报错：
// main.go:11: invalid operation: <-sendch (receive from send-only type chan<- int)

package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)
}

// 在上述程序的第 10 行，我们创建了一个双向信道 cha1。
// 在第 11 行 cha1 作为参数传递给了 sendData 协程。
// 在第 5 行，函数 sendData 里的参数 sendch chan<- int 把 cha1 转换为一个唯送信道。
// 于是该信道在 sendData 协程里是一个唯送信道，而在 Go 主协程里是一个双向信道。该程序最终打印输出 10。
