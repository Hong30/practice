package main

import (
	"fmt"
)

func main() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}
}

// map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
// personSalary 是 nil，因此需要使用 make 方法初始化，程序将输出 map is nil. Going to make one.。
