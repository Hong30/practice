// package main

// import (
// 	"fmt"
// )

// type Employee struct {
// 	firstName, lastName string
// 	age, salary         int
// }

// func main() {
// 	var emp4 Employee //zero valued structure
// 	fmt.Println("Employee 4", emp4)
// }
// 该程序定义了 emp4，却没有初始化任何值。因此 firstName 和 lastName 赋值为 string 的零值（""）。
// 而 age 和 salary 赋值为 int 的零值（0）。

// 当然还可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值。
package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)
}

// 初始化了 firstName 和 lastName，而 age 和 salary 没有进行初始化。
// 因此 age 和 salary 赋值为零值。
