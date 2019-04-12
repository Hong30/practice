package main

import (
	"fmt"
)

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
}

// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = (length + width) * 2
// 	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
// }
