package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p.area() //通过指针调用值接收器
}

// 第 12 行的函数 func area(r rectangle) 接受一个值参数，方法 func (r rectangle) area() 接受一个值接收器。
// 在第 25 行，我们通过值参数 area(r) 来调用 area 这个函数，这是合法的。
// 同样，我们使用值接收器来调用 area 方法 r.area()，这也是合法的。
// 在第 28 行，我们创建了一个指向 r 的指针 p。
// 如果我们试图把这个指针传递到只能接受一个值参数的函数 area，编译器将会报错。所以我把代码的第 33 行注释了。
// 如果你把这行的代码注释去掉，编译器将会抛出错误 compilation error,
// cannot use p (type *rectangle) as type rectangle in argument to area.。这将会按预期抛出错误。
// 现在到了棘手的部分了，在第35行的代码 p.area() 使用指针接收器 p 调用了只接受一个值接收器的方法 area。这是完全有效的。
// 原因是当 area 有一个值接收器时，为了方便Go语言把 p.area() 解释为 (*p).area()。
