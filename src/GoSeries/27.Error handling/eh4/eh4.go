package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

// 在上述程序里，我们查询了模式为 [ 的文件，然而这个模式写的不正确。
// 我们检查了该错误是否为 nil。为了获取该错误的更多信息，我们在第 10 行将 error 直接与 filepath.ErrBadPattern 相比较。
// 如果该条件满足，那么该错误就是由模式错误导致的。

// 该程序会输出：

// syntax error in pattern
