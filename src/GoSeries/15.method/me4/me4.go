package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

/*
使用值接收器的方法。
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

// 在上面的程序中，changeName 方法有一个值接收器 (e Employee)，
// 而 changeAge 方法有一个指针接收器 (e *Employee)。
// 在 changeName 方法中对 Employee 结构体的字段 name 所做的改变对调用者是不可见的，
// 因此程序在调用 e.changeName("Michael Andrew") 这个方法的前后打印出相同的名字。
// 由于 changeAge 方法是使用指针 (e *Employee) 接收器的，
// 所以在调用 (&e).changeAge(51) 方法对 age 字段做出的改变对调用者将是可见的。
