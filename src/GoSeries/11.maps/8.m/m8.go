package main

import (
	"fmt"
)

func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

}

// 第 14 行，personSalary 被赋值给 newPersonSalary。下一行 ，newPersonSalary 中 mike 的薪资变成了 18000 。
// personSalary 中 Mike 的薪资也会变成 18000。
