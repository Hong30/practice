package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}

// 在上面的程序中，因为 NumField 方法只能在结构体上使用，
// 我们在第 14 行首先检查了 q 的类别是 struct。程序的其他代码很容易看懂，不作解释。该程序会输出：

// Number of fields 2
// Field:0 type:reflect.Value value:456
// Field:1 type:reflect.Value value:56
