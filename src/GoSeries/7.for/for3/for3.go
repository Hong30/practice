package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

// if i%2==0 会判断 i 除以 2 的余数是不是 0，如果是 0，这个数字就是偶数然后执行 continue 语句，从而控制程序进入下一个循环。
// 因此在 continue 后面的打印语句不会被调用而程序会进入一下个循环。
