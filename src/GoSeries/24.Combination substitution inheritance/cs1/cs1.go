package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// 在上面的代码片段中，我们创建了一个 author 结构体，author 的字段有 firstname、lastname 和 bio。
// 我们还添加了一个 fullName() 方法，其中 author 作为接收者类型，该方法返回了作者的全名。

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Bio: ", p.author.bio)
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}

// post 结构体的字段有 title 和 content。
// 它还有一个嵌套的匿名字段 author。该字段指定 author 组成了 post 结构体。
// 现在 post 可以访问 author 结构体的所有字段和方法。
// 我们同样给 post 结构体添加了 details() 方法，用于打印标题、内容和作者的全名与简介。
// 一旦结构体内嵌套了一个结构体字段，Go 可以使我们访问其嵌套的字段，好像这些字段属于外部结构体一样。
// 所以上面第 11 行的 p.author.fullName() 可以替换为 p.fullName()。

// 在上面程序中，main 函数在第 31 行新建了一个 author 结构体变量。
// 而在第 36 行，我们通过嵌套 author1 来创建一个 post。
// 该程序输出：
// Title:  Inheritance in Go
// Content:  Go supports composition instead of inheritance
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast
