package main

import "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
}

// 接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
// 上面程序里的 d1 等于 nil，程序会输出：
// d1 is nil and has type <nil> value <nil>

// package main

// type Describer interface {
//     Describe()
// }

// func main() {
//     var d1 Describer
//     d1.Describe()
// }

// 在上述程序中，d1 等于 nil，程序产生运行时错误
// panic： panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xc8527] 。
