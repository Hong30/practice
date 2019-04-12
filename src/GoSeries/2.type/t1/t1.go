package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}

// a 赋值为 true，b 赋值为 false。
// c 赋值为 a && b。仅当 a 和 b 都为 true 时，操作符 && 才返回 true。因此，在这里 c 为 false。
// 当 a 或者 b 为 true 时，操作符 || 返回 true。在这里，由于 a 为 true，因此 d 也为 true。
