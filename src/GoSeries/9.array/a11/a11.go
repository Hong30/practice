package main

import (
	"fmt"
)

func main() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}

// 当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中。
// 在 9 行中，numa [:] 缺少开始和结束值。
// 开始和结束的默认值分别为 0 和 len (numa)。
// 两个切片 nums1 和 nums2 共享相同的数组。
