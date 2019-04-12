package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

// 在第 22 行延迟了一个方法调用。而其他的代码很直观。
// 该程序输出：
// Welcome John Smith
