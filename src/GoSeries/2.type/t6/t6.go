// package main

// import (
//     "fmt"
// )

// func main() {
//     i := 55      //int
//     j := 67.8    //float64
//     sum := i + j //不允许 int + float64
//     fmt.Println(sum)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	i := 55           //int
// 	j := 67.8         //float64
// 	sum := i + int(j) //j is converted to int
// 	fmt.Println(sum)
// }

package main

import (
	"fmt"
)

func main() {
	i := 10
	var j float64 = float64(i) // 若没有显式转换，该语句会报错
	// i 转换为 float64 类型，接下来赋值给 j
	fmt.Println("j", j)
}
