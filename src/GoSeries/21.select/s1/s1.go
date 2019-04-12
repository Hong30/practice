package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

// 在上面程序里，server1 函数（第 8 行）休眠了 6 秒，接着将文本 from server1 写入信道 ch。
// 而 server2 函数（第 12 行）休眠了 3 秒，然后把 from server2 写入了信道 ch。
// 而 main 函数在第 20 行和第 21 行，分别调用了 server1 和 server2 两个 Go 协程。
// 在第 22 行，程序运行到了 select 语句。select 会一直发生阻塞，除非其中有 case 准备就绪。
// 在上述程序里，server1 协程会在 6 秒之后写入 output1 信道，而server2 协程在 3 秒之后就写入了 output2 信道。
// 因此 select 语句会阻塞 3 秒钟，等着 server2 向 output2 信道写入数据。3 秒钟过后，程序会输出：
// from server2
// 然后程序终止。
