package main

import (
	"fmt"
)

func main() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}

// 在上述程序的第 10 行，food 是通过 append(veggies, fruits...) 创建。
// 程序的输出为 food: [potatoes tomatoes brinjal oranges apples]。
