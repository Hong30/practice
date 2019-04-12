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
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}

// 在上面的程序中，第 13 行的 createQuery 函数接收 interface{} 作为参数。
// 在第 14 行，reflect.TypeOf 接收了参数 interface{}，返回了reflect.Type，它包含了传入的 interface{} 参数的具体类型。
// 同样地，在第 15 行，reflect.ValueOf 函数接收参数 interface{}，并返回了 reflect.Value，它包含了传来的 interface{} 的具体值。

// 上述程序会打印：

// Type  main.order
// Value  {456 56}
