package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
func main() {
	findType("Naveen")
	findType(77)
	findType(89.98)
}

// 在上述程序的第 8 行，switch i.(type) 表示一个类型选择。
// 每个 case 语句都把 i 的具体类型和一个指定类型进行了比较。
// 如果 case 匹配成功，会打印出相应的语句。
//第 20 行中的 89.98 的类型是 float64，没有在 case 上匹配成功，因此最后一行打印了 Unknown type
