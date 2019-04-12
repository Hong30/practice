package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

// 你认为这段代码将输出什么呢？如果你认为它输出 [Go world] 。恭喜你！你已经理解了可变参数函数和切片。
// 如果你猜错了，那也不要紧，让我来解释下为什么会有这样的输出。
// 在第 13 行，我们使用了语法糖 ... 并且将切片作为可变参数传入 change 函数。
// 正如前面我们所讨论的，如果使用了 ... ，welcome 切片本身会作为参数直接传入，
// 不需要再创建一个新的切片。这样参数 welcome 将作为参数传入 change 函数
// 在 change 函数中，切片的第一个元素被替换成 Go
