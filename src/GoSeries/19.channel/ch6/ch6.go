// package main

// import (
// 	"fmt"
// )

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }
// func main() {
// 	ch := make(chan int)
// 	go producer(ch)
// 	for {
// 		v, ok := <-ch
// 		if ok == false {
// 			break
// 		}
// 		fmt.Println("Received ", v, ok)
// 	}
// }

//  在上述的程序中，producer 协程会从 0 到 9 写入信道 chn1，然后关闭该信道。
//  主函数有一个无限的 for 循环（第 16 行），使用变量 ok（第 18 行）检查信道是否已经关闭。
//  如果 ok 等于 false，说明信道已经关闭，于是退出 for 循环。如果 ok 等于 true，会打印出接收到的值和 ok 的值。

// package main

// import (
//     "fmt"
// )

// func producer(chnl chan int) {
//     for i := 0; i < 10; i++ {
//         chnl <- i
//     }
//     close(chnl)
// }
// func main() {
//     ch := make(chan int)
//     go producer(ch)
//     for v := range ch {
//         fmt.Println("Received ",v)
//     }
// }

// for range 循环从信道 ch 接收数据，直到该信道关闭。
// 一旦关闭了 ch，循环会自动结束。

// 我们可以使用 for range 循环，重写信道的另一个示例这一节里面的代码，提高代码的可重用性。
// 如果你仔细观察这段代码，会发现获得一个数里的每位数的代码在 calcSquares 和 calcCubes 两个函数内重复了。
// 我们将把这段代码抽离出来，放在一个单独的函数里，然后并发地调用它。

package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

// 上述程序里的 digits 函数，包含了获取一个数的每位数的逻辑，并且 calcSquares 和 calcCubes 两个函数并发地调用了 digits。
// 当计算完数字里面的每一位数时，第 13 行就会关闭信道。
// calcSquares 和 calcCubes 两个协程使用 for range 循环分别监听了它们的信道，直到该信道关闭。
// 程序的其他地方不变，该程序同样会输出：
// Final output 1536
