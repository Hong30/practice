package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
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

// 在第 7 行，函数 calcSquares 计算一个数每位的平方和，并把结果发送给信道 squareop。
// 与此类似，在第 17 行函数 calcCubes 计算一个数每位的立方和，并把结果发送给信道 cubop。
// 这两个函数分别在单独的协程里运行（第 31 行和第 32 行），每个函数都有传递信道的参数，以便写入数据。
// Go 主协程会在第 33 行等待两个信道传来的数据。
// 一旦从两个信道接收完数据，数据就会存储在变量 squares 和 cubes 里，然后计算并打印出最后结果。
