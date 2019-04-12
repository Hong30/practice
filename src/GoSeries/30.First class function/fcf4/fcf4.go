package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}

// 写一个 filter 函数。该函数接收一个 students 切片和一个函数作为参数，这个函数会计算一个学生是否满足筛选条件。
// filter 的第二个参数是一个函数。这个函数接收 student 参数，返回一个 bool 值。
// 这个函数计算了某一学生是否满足筛选条件。
// 我们在第 3 行遍历了 student 切片，将每个学生作为参数传递给了函数 f。
// 如果该函数返回 true，就表示该学生通过了筛选条件，接着将该学生添加到了结果切片 r 中。
// 你可能会很困惑这个函数的实际用途，等我们完成程序你就知道了。
// 在 main 函数中，我们首先创建了两个学生 s1 和 s2，并将他们添加到了切片 s。
// 现在假设我们想要查询所有成绩为 B 的学生。
// 为了实现这样的功能，我们传递了一个检查学生成绩是否为 B 的函数，如果是，该函数会返回 true。我们把这个函数作为参数传递给了 filter 函数（第 38 行）。

// 上述程序会输出：

// [{Samuel Johnson B USA}]
