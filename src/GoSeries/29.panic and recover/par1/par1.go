package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
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
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

// 上面的程序很简单，会打印一个人的全名。
// 第 7 行的 fullName 函数会打印出一个人的全名。
// 该函数在第 8 行和第 11 行分别检查了 firstName 和 lastName 的指针是否为 nil。
// 如果是 nil，fullName 函数会调用含有不同的错误信息的 panic。当程序终止时，会打印出该错误信息。

// 运行该程序，会有如下输出：

// panic: runtime error: last name cannot be nil

// goroutine 1 [running]:
// main.fullName(0x1040c128, 0x0)
//     /tmp/sandbox135038844/main.go:12 +0x120
// main.main()
// 	/tmp/sandbox135038844/main.go:20 +0x80

// 来分析这个输出，理解一下 panic 是如何工作的，并且思考当程序发生 panic 时，会怎样打印堆栈跟踪。
// 在第 19 行，我们将 Elon 赋值给了 firstName。
// 在第 20 行，我们调用了 fullName 函数，其中 lastName 等于 nil。
// 因此，满足了第 11 行的条件，程序发生 panic。当出现了 panic 时，程序就会终止运行，打印出传入 panic 的参数，接着打印出堆栈跟踪。
// 因此，第 14 行和第 15 行的代码并不会在发生 panic 之后执行。
// 程序首先会打印出传入 panic 函数的信息：
// panic: runtime error: last name cannot be empty
// 接着打印出堆栈跟踪。
// 程序在 fullName 函数的第 12 行发生 panic，因此，首先会打印出如下所示的输出。
// main.fullName(0x1040c128, 0x0)
// 	/tmp/sandbox135038844/main.go:12 +0x120
// 接着会打印出堆栈的下一项。
// 在本例中，堆栈跟踪中的下一项是第 20 行（因为发生 panic 的 fullName 调用就在这一行），因此接下来会打印出：
// 	main.main()
// 		/tmp/sandbox135038844/main.go:20 +0x80
// 现在我们已经到达了导致 panic 的顶层函数，这里没有更多的层级，因此结束打印。

// package main

// import (
//     "fmt"
// )

// func fullName(firstName *string, lastName *string) {
//     defer fmt.Println("deferred call in fullName")
//     if firstName == nil {
//         panic("runtime error: first name cannot be nil")
//     }
//     if lastName == nil {
//         panic("runtime error: last name cannot be nil")
//     }
//     fmt.Printf("%s %s\n", *firstName, *lastName)
//     fmt.Println("returned normally from fullName")
// }

// func main() {
//     defer fmt.Println("deferred call in main")
//     firstName := "Elon"
//     fullName(&firstName, nil)
//     fmt.Println("returned normally from main")
// }

// 上述代码中，我们只修改了两处，分别在第 8 行和第 20 行添加了延迟函数的调用。

// 该函数会打印：

// This program prints,

// deferred call in fullName
// deferred call in main
// panic: runtime error: last name cannot be nil

// goroutine 1 [running]:
// main.fullName(0x1042bf90, 0x0)
//     /tmp/sandbox060731990/main.go:13 +0x280
// main.main()
//     /tmp/sandbox060731990/main.go:22 +0xc0
// 当程序在第 13 行发生 panic 时，首先执行了延迟函数，接着控制返回到函数调用方，调用方的延迟函数继续运行，直到到达顶层调用函数。

// 在我们的例子中，首先执行 fullName 函数中的 defer 语句（第 8 行）。程序打印出：

// deferred call in fullName
// 接着程序返回到 main 函数，执行了 main 函数的延迟调用，因此会输出：

// deferred call in main
// 现在程序控制到达了顶层函数，因此该函数会打印出 panic 信息，然后是堆栈跟踪，最后终止程序。
