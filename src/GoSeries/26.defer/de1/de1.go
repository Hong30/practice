package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}

// 上面的程序很简单，就是找出一个给定切片的最大值。
// largest 函数接收一个 int 类型的切片作为参数，然后打印出该切片中的最大值。
// largest 函数的第一行的语句为 defer finished()。
// 这表示在 finished() 函数将要返回之前，会调用 finished() 函数。

// 运行该程序，你会看到有如下输出：

// Started finding largest
// Largest number in [78 109 2 563 300] is 563
// Finished finding largest

// largest 函数开始执行后，会打印上面的两行输出。
// 而就在 largest 将要返回的时候，又调用了我们的延迟函数（Deferred Function），打印出 Finished finding largest 的文本。
