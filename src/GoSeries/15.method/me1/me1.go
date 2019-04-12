package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
  displaySalary() 方法将 Employee 做为接收器类型
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法
}

// 在上面程序的第 16 行，我们在 Employee 结构体类型上创建了一个 displaySalary 方法。
// displaySalary()方法在方法的内部访问了接收器 e Employee。
// 在第 17 行，我们使用接收器 e，并打印 employee 的 name、currency 和 salary 这 3 个字段。
// 在第 26 行，我们调用了方法 emp1.displaySalary()。
