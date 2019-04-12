package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
}

// 条件语句会检查 i 是否小于 10。
// 如果条件成立，i 就会被打印出来，否则循环就会终止。
// 循环语句会在每一次循环完成后自增 1。
// 一旦 i 变得比 10 要大，循环中止
