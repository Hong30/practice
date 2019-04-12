package main

import (
	"fmt"
)

// func functionname(parametername type) returntype {
//     // 函数体（具体实现的功能）
// }

// func functionname() {
//     // 译注: 表示这个函数不需要输入参数，且没有返回值
// }

// func calculateBill(price int, no int) int {
//     var totalPrice = price * no // 商品总价 = 商品单价 * 数量
//     return totalPrice // 返回总价
// }

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
func main() {
	price, no := 90, 6 // 定义 price 和 no,默认类型为 int
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice) // 打印到控制台上
}
