// package main

// import (
//     "fmt"
//     "sync"
// )

// type rect struct {
//     length int
//     width  int
// }

// func (r rect) area(wg *sync.WaitGroup) {
//     if r.length < 0 {
//         fmt.Printf("rect %v's length should be greater than zero\n", r)
//         wg.Done()
//         return
//     }
//     if r.width < 0 {
//         fmt.Printf("rect %v's width should be greater than zero\n", r)
//         wg.Done()
//         return
//     }
//     area := r.length * r.width
//     fmt.Printf("rect %v's area %d\n", r, area)
//     wg.Done()
// }

// func main() {
//     var wg sync.WaitGroup
//     r1 := rect{-67, 89}
//     r2 := rect{5, -67}
//     r3 := rect{8, 9}
//     rects := []rect{r1, r2, r3}
//     for _, v := range rects {
//         wg.Add(1)
//         go v.area(&wg)
//     }
//     wg.Wait()
//     fmt.Println("All go routines finished executing")
// }

// 在上面的程序里，我们在第 8 行创建了 rect 结构体，并在第 13 行创建了 rect 的方法 area，计算出矩形的面积。
// area 检查了矩形的长宽是否小于零。
// 如果矩形的长宽小于零，它会打印出对应的提示信息，而如果大于零，它会打印出矩形的面积。
// main 函数创建了 3 个 rect 类型的变量：r1、r2 和 r3。在第 34 行，我们把这 3 个变量添加到了 rects 切片里。
// 该切片接着使用 for range 循环遍历，把 area 方法作为一个并发的 Go 协程进行调用（第 37 行）。
// 我们用 WaitGroup wg 来确保 main 函数在其他协程执行完毕之后，才会结束执行。
// WaitGroup 作为参数传递给 area 方法后，在第 16 行、第 21 行和第 26 行通知 main 函数，表示现在协程已经完成所有任务。
// 如果你仔细观察，会发现 wg.Done() 只在 area 函数返回的时候才会调用。
// wg.Done() 应该在 area 将要返回之前调用，并且与代码流的路径（Path）无关，因此我们可以只调用一次 defer，来有效地替换掉 wg.Done() 的多次调用。
// 我们来用 defer 来重写上面的代码。
// 在下面的代码中，我们移除了原先程序中的 3 个 wg.Done 的调用，而是用一个单独的 defer wg.Done() 来取代它（第 14 行）。

package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// 该程序会输出：

// rect {8 9}'s area 72
// rect {-67 89}'s length should be greater than zero
// rect {5 -67}'s width should be greater than zero
// All go routines finished executing
// 在上面的程序中，使用 defer 还有一个好处。
// 假设我们使用 if 条件语句，又给 area 方法添加了一条返回路径（Return Path）。
// 如果没有使用 defer 来调用 wg.Done()，我们就得很小心了，确保在这条新添的返回路径里调用了 wg.Done()。
// 由于现在我们延迟调用了 wg.Done()，因此无需再为这条新的返回路径添加 wg.Done() 了。
