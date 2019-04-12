package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
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

// 在上面的程序中，我们使用 Errorf（第 10 行）打印了发生错误的半径。

// 程序运行后会输出：

// Area calculation failed, radius -20.00 is less than zero
