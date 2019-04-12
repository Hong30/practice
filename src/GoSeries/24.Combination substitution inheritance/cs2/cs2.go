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

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
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
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}

// 在上面的主函数中，我们创建了一个作者 author1，以及三个帖子 post1、post2 和 post3。
// 我们最后通过嵌套三个帖子，在第 62 行创建了网站 w，并在下一行显示内容。
// 程序会输出：
// Contents of Website

// Title:  Inheritance in Go
// Content:  Go supports composition instead of inheritance
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast

// Title:  Struct instead of Classes in Go
// Content:  Go does not support classes but methods can be added to structs
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast

// Title:  Concurrency
// Content:  Go is a concurrent language and not a parallel one
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast
