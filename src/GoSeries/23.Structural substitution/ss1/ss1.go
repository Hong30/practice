// package employee

// import (
// 	"fmt"
// )

// type Employee struct {
// 	FirstName   string
// 	LastName    string
// 	TotalLeaves int
// 	LeavesTaken int
// }

// func (e Employee) LeavesRemaining() {
// 	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
// }

// 在上述程序里，第 1 行指定了该文件属于 employee 包。
// 而第 7 行声明了一个 Employee 结构体。在第 14 行，结构体 Employee 添加了一个名为 LeavesRemaining 的方法。
// 该方法会计算和显示员工的剩余休假数。于是现在我们有了一个结构体，并绑定了结构体的方法，这与类很相似。
// 接着在 oop 文件夹里创建一个文件，命名为 main.go。
// 现在目录结构如下所示：
// workspacepath -> oop -> employee -> employee.go
// workspacepath -> oop -> main.go

package employee

import (
	"fmt"
)

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

// 我们进行了一些重要的修改。我们把 Employee 结构体的首字母改为小写 e，也就是将 type Employee struct 改为了 type employee struct。
// 通过这种方法，我们把 employee 结构体变为了不可引用的，防止其他包对它的访问。
// 除非有特殊需求，否则也要隐藏所有不可引用的结构体的所有字段，这是 Go 的最佳实践。
// 由于我们不会在外部包需要 employee 的字段，因此我们也让这些字段无法引用。
// 同样，我们还修改了 LeavesRemaining() 的方法。
// 现在由于 employee 不可引用，因此不能在其他包内直接创建 Employee 类型的变量。
// 于是我们在第 14 行提供了一个可引用的 New 函数，该函数接收必要的参数，返回一个新创建的 employee 结构体变量。
// 这个程序还需要一些必要的修改，但现在先运行这个程序，理解一下当前的修改。如果运行当前程序，编译器会报错，如下所示：
// go/src/constructor/main.go:6: undefined: employee.Employee
