package main

import (
	"fmt"
)

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
}

// fruitslice 是从 fruitarray 的索引 1 和 2 创建的。 因此，fruitlice 的长度为 2。
// fruitarray 的长度是 7。fruiteslice 是从 fruitarray 的索引 1 创建的。
// 因此, fruitslice 的容量是从 fruitarray 索引为 1 开始，也就是说从 orange 开始，该值是 6。
// 因此, fruitslice 的容量为 6。该程序输出切片的 长度为 2 容量为 6 。
