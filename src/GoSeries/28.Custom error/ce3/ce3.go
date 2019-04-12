package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}

// 在上面的程序中，circleArea（第 17 行）用于计算圆的面积。
// 该函数首先检查半径是否小于零，如果小于零，它会通过错误半径和对应错误信息，
// 创建一个 areaError 类型的值，然后返回 areaError 值的地址，与此同时 area 等于 0（第 19 行）。
// 于是我们提供了更多的错误信息（即导致错误的半径），我们使用了自定义错误的结构体字段来定义它。
// 如果半径是非负数，该函数会在第 21 行计算并返回面积，同时错误值为 nil。
// 在 main 函数的 26 行，我们试图计算半径为 -20 的圆的面积。
// 由于半径小于零，因此会导致一个错误。
// 我们在第 27 行检查了错误是否为 nil，并在下一行断言了 *areaError 类型。
// 如果错误是 *areaError 类型，我们就可以用 err.radius 来获取错误的半径（第 29 行），打印出自定义错误的消息，
// 最后程序返回退出。
// 如果断言错误，我们就在第 32 行打印该错误，并返回。
// 如果没有发生错误，在第 35 行会打印出面积。

// 该程序会输出：

// Radius -20.00 is less than zero
