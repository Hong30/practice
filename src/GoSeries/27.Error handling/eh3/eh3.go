package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

// 在上述程序中，我们在第 9 行，试图获取 golangbot123.com（无效的域名） 的 ip。
// 在第 10 行，我们通过 *net.DNSError 的类型断言，获取到了错误的底层值。
// 接下来的第 11 行和第 13 行，我们分别检查了该错误是由超时引起的，还是一个临时性错误。
// 在本例中，我们的错误既不是临时性错误，也不是由超时引起的，

// 因此该程序输出：

// generic error:  lookup golangbot123.com: no such host
