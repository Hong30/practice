package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}

// 在第 9 行，通过使用 _ 空白标识符，我忽略了 Glob 函数返回的错误。我在第 10 行简单打印了所有匹配的文件。

// 该程序会输出：

// matched files []
