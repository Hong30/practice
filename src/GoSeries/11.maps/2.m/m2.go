package main

import (
	"fmt"
)

func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
}

// 可以在声明的时候初始化 map。
// package main

// import (
//     "fmt"
// )

// func main() {
//     personSalary := map[string]int {
//         "steve": 12000,
//         "jamie": 15000,
//     }
//     personSalary["mike"] = 9000
//     fmt.Println("personSalary map contents:", personSalary)
// }

// 程序声明了 personSalary，并在声明的同时添加两个元素。之后又添加了键 mike。
