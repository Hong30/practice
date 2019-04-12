package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

// 在上面程序里，我们在第 18 行和第 19 行分别调用了 server1 和 server2 两个 Go 协程。
// 接下来，主程序休眠了 1 秒钟（第 20 行）。
// 当程序控制到达第 21 行的 select 语句时，server1 已经把 from server1 写到了 output1 信道上，
// 而 server2 也同样把 from server2 写到了 output2 信道上。
// 因此这个 select 语句中的两种情况都准备好执行了。
// 如果你运行这个程序很多次的话，输出会是 from server1 或者 from server2，这会根据随机选取的结果而变化。
