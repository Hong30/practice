package main

import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

// 检查半径是否小于零（第 10 行）。
// 如果半径小于零，我们会返回等于 0 的面积，以及相应的错误信息。
// 如果半径大于零，则会计算出面积，并返回值为 nil 的错误（第 13 行）。
// 在 main 函数里，我们在第 19 行检查错误是否等于 nil。
// 如果不是 nil，我们会打印出错误并返回，否则我们会打印出圆的面积。
// 在我们的程序中，半径小于零，

// 因此打印出：

// Area calculation failed, radius is less than zero
