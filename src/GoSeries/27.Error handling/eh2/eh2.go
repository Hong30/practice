package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// 在上面的程序里，我们在第 10 行使用了类型断言（Type Assertion）来获取 error 接口的底层值（Underlying Value）。
// 接下来在第 11 行，我们使用 err.Path 来打印该路径。

// 该程序会输出：

// File at path /test.txt failed to open
