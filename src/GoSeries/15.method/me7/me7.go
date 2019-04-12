package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	p := &r //pointer to r
	perimeter(p)
	p.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	r.perimeter() //使用值来调用指针接收器
}

// 在上面程序的第 12 行，定义了一个接受指针参数的函数 perimeter。
// 第 17 行定义了一个有一个指针接收器的方法。
// 在第 27 行，我们调用 perimeter 函数时传入了一个指针参数。
// 在第 28 行，我们通过指针接收器调用了 perimeter 方法。所有一切看起来都这么完美。
// 在被注释掉的第 33 行，我们尝试通过传入值参数 r 调用函数 perimeter。
// 这是不被允许的，因为函数的指针参数不接受值参数。
// 如果你把这行的代码注释去掉并把程序运行起来，
// 编译器将会抛出错误 main.go:33: cannot use r (type rectangle) as type *rectangle in argument to perimeter.。
// 在第 35 行，我们通过值接收器 r 来调用有指针接收器的方法 perimeter。这是被允许的，
// 为了方便Go语言把代码 r.perimeter() 解释为 (&r).perimeter()。
