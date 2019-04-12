package main

import (
	"fmt"
)

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

// 切片自己不拥有任何数据。它只是底层数组的一种表示。
// 对切片所做的任何修改都会反映在底层数组中。
// 数组索引 2,3,4 创建一个切片 dslice。for 循环将这些索引中的值逐个递增。
// 当我们使用 for 循环打印数组时，我们可以看到对切片的更改反映在数组中。
