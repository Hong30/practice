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
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"])
}

// 程序返回 joe 的薪资是 0。personSalary 中不包含 joe 的情况下我们不会获取到任何运行时错误。
