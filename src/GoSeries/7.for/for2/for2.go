package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

// 在循环过程中 i 的值会被判断。
// 如果 i 的值大于 5 然后 break 语句就会执行，循环就会被终止。
// 打印语句会在 for 循环结束后执行，
