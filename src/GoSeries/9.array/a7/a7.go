package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

// for i, v := range a 利用的是 for 循环 range 方式。
// 它将返回索引和该索引处的值。
// 我们打印这些值，并计算数组 a 中所有元素的总和。
