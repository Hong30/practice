package main

import (
	"fmt"
)

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        // re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

// 第 11 行中，fruitslice 的容量是重置的。
