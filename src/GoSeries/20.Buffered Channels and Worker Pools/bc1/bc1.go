package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 在上面程序里的第 9 行，我们创建了一个缓冲信道，其容量为 2。
// 由于该信道的容量为 2，因此可向它写入两个字符串，而且不会发生阻塞。在第 10 行和第 11 行，
// 我们向信道写入两个字符串，该信道并没有发生阻塞。
// 我们又在第 12 行和第 13 行分别读取了这两个字符串。该程序输出：
// naveen
// paul
