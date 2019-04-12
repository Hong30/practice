package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())
}

// Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
// 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
// 假设我们有一个 Square 和 Circle 结构体。可以在 Square 和 Circle 上分别定义一个 Area 方法。
