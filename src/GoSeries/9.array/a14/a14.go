package main

import (
	"fmt"
)

func main() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}

// 使用 make 创建切片时默认情况下这些值为零。上述程序的输出为 [0 0 0 0 0]。
