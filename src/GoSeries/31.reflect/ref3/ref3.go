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
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}

// 上述程序会输出：

// Type  main.order
// Kind  struct
// 我想你应该很清楚两者的区别了。
// Type 表示 interface{} 的实际类型（在这里是 main.Order)，而 Kind 表示该类型的特定类别（在这里是 struct）。
