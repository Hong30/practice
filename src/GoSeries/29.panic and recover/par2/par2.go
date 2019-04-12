package main

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

// 在第 7 行，recoverName() 函数调用了 recover()，返回了调用 panic 的传参。
// 在这里，我们只是打印出 recover 的返回值（第 8 行）。
// 在 fullName 函数内，我们在第 14 行延迟调用了 recoverNames()。

// 当 fullName 发生 panic 时，会调用延迟函数 recoverName()，它使用了 recover() 来停止 panic 续发事件。

// 该程序会输出：

// recovered from  runtime error: last name cannot be nil
// returned normally from main
// deferred call in main

// 当程序在第 19 行发生 panic 时，会调用延迟函数 recoverName，它反过来会调用 recover() 来重新获得 panic 协程的控制。
// 第 8 行调用了 recover，返回了 panic 的传参，因此会打印：

// recovered from  runtime error: last name cannot be nil

// 在执行完 recover() 之后，panic 会停止，程序控制返回到调用方（在这里就是 main 函数），
// 程序在发生 panic 之后，从第 29 行开始会继续正常地运行。
// 程序会打印 returned normally from main，之后是 deferred call in main。
