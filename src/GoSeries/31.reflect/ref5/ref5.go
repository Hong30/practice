package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

}

// 在第 22 行，我们首先检查了传来的参数是否是一个结构体。
// 在第 23 行，我们使用了 Name() 方法，从该结构体的 reflect.Type 获取了结构体的名字。
// 接下来一行，我们用 t 来创建查询。

// 在第 28 行，case 语句 检查了当前字段是否为 reflect.Int，
// 如果是的话，我们会取到该字段的值，并使用 Int() 方法转换为 int64。if else 语句用于处理边界情况。
// 请添加日志来理解为什么需要它。在第 34 行，我们用来相同的逻辑来取到 string。

// 我们还作了额外的检查，以防止 createQuery 函数传入不支持的类型时，程序发生崩溃。
// 程序的其他代码是自解释性的。我建议你在合适的地方添加日志，检查输出，来更好地理解这个程序。

// 该程序会输出：

// insert into order values(456, 56)
// insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
// unsupported type
// 至于向输出的查询中添加字段名，我们把它留给读者作为练习。请尝试着修改程序，打印出以下格式的查询。

// insert into order(ordId, customerId) values(456, 56)
