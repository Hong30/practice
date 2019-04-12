package main

import (
	"fmt"
)

func main() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
}

// cars 的容量最初是 3。在第 10 行，我们给 cars 添加了一个新的元素，
// 并把 append(cars, "Toyota") 返回的切片赋值给 cars。现在 cars 的容量翻了一番，变成了 6。
