// package main

// func main() {
//     ch := make(chan string)
//     select {
//     case <-ch:
//     }
// }

// 上面的程序中，我们在第 4 行创建了一个信道 ch。
// 我们在 select 内部（第 6 行），试图读取信道 ch。
// 由于没有 Go 协程向该信道写入数据，因此 select 语句会一直阻塞，导致死锁。该程序会触发运行时 panic，报错信息如下：
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan receive]:
// main.main()
// 	/tmp/sandbox416567824/main.go:6 +0x80

// 如果存在默认情况，就不会发生死锁，因为在没有其他 case 准备就绪时，会执行默认情况。我们用默认情况重写后，程序如下：
// package main

// import "fmt"

// func main() {
//     ch := make(chan string)
//     select {
//     case <-ch:
//     default:
//         fmt.Println("default case executed")
//     }
// }

// 以上程序会输出：
// default case executed

// 如果 select 只含有值为 nil 的信道，也同样会执行默认情况。

package main

import "fmt"

func main() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")

	}
}

// 在上面程序中，ch 等于 nil，而我们试图在 select 中读取 ch（第 8 行）。
// 如果没有默认情况，select 会一直阻塞，导致死锁。
// 由于我们在 select 内部加入了默认情况，程序会执行它，并输出：
// default case executed
