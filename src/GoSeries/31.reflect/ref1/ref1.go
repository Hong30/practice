// package main

// import (
//     "fmt"
// )

// func main() {
//     i := 10
//     fmt.Printf("%d %T", i, i)
// }

// 在上面的程序中，i 的类型在编译时就知道了，然后我们在下一行打印出 i。

// package main

// import (
//     "fmt"
// )

// type order struct {
//     ordId      int
//     customerId int
// }

// func main() {
//     o := order{
//         ordId:      1234,
//         customerId: 567,
//     }
//     fmt.Println(o)
// }

// 在上面的程序中，我们需要编写一个函数，接收结构体变量 o 作为参数，返回下面的 SQL 插入查询。

// insert into order values(1234, 567)

package main

import (
	"fmt"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func main() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o))
}

// 在第 12 行，createQuery 函数用 o 的两个字段（ordId 和 customerId），创建了插入查询。

// 该程序会输出：

// insert into order values(1234, 567)

// package main

// type order struct {
//     ordId      int
//     customerId int
// }

// type employee struct {
//     name string
//     id int
//     address string
//     salary int
//     country string
// }

// func createQuery(q interface{}) string {
// }

// func main() {

// }

// 我们的目标就是完成 createQuery 函数（上述程序中的第 16 行），它可以接收任何结构体作为参数，根据结构体的字段创建插入查询。

// 例如，如果我们传入下面的结构体：

// o := order {
//     ordId: 1234,
//     customerId: 567
// }
// createQuery 函数应该返回：

// insert into order values (1234, 567)

// 类似地，如果我们传入：

//  e := employee {
//         name: "Naveen",
//         id: 565,
//         address: "Science Park Road, Singapore",
//         salary: 90000,
//         country: "Singapore",
//     }
// 该函数会返回：

// insert into employee values("Naveen", 565, "Science Park Road, Singapore", 90000, "Singapore")
// 由于 createQuery 函数应该适用于任何结构体，因此它接收 interface{} 作为参数。
// 为了简单起见，我们只处理包含 string 和 int 类型字段的结构体，但可以扩展为包含任何类型的字段。

// createQuery 函数应该适用于所有的结构体。
// 因此，要编写这个函数，就必须在运行时检查传递过来的结构体参数的类型，找到结构体字段，接着创建查询。
// 这时就需要用到反射了。在本教程的下一步，我们将会学习如何使用 reflect 包来实现它。
