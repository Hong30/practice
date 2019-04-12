package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

// 在第 19 行，我们延迟了文件的关闭操作。

// 在上面程序的第 24 行，我们新建了一个缓冲读取器（buffered reader）。
// 在下一行，我们创建了长度和容量为 3 的字节切片，程序会把文件的字节读取到切片中。

// 第 27 行的 Read 方法会读取 len(b) 个字节（达到 3 字节），并返回所读取的字节数。
// 当到达文件最后时，它会返回一个 EOF 错误。程序的其他地方比较简单，不做解释。

// 如果我们使用下面命令来运行程序：

// $ go install filehandling
// $ wrkspacepath/bin/filehandling -fpath=/path-of-file/test.txt
// 会得到以下输出：

// Hel
// lo
// Wor
// ld.
//  We
// lco
// me
// to
// fil
// e h
// and
// lin
// g i
// n G
// o.
// Error reading file: EOF

// 将 filehandling.go 替换为以下内容。

// package main

// import (
//     "bufio"
//     "flag"
//     "fmt"
//     "log"
//     "os"
// )

// func main() {
//     fptr := flag.String("fpath", "test.txt", "file path to read from")
//     flag.Parse()

//     f, err := os.Open(*fptr)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer func() {
//         if err = f.Close(); err != nil {
//         log.Fatal(err)
//     }
//     }()
//     s := bufio.NewScanner(f)
//     for s.Scan() {
//         fmt.Println(s.Text())
//     }
//     err = s.Err()
//     if err != nil {
//         log.Fatal(err)
//     }
// }

// 在上述程序的第 15 行，我们用命令行标记传入的路径，打开文件。
// 在第 24 行，我们用文件创建了一个新的 scanner。
// 第 25 行的 Scan() 方法读取文件的下一行，如果可以读取，就可以使用 Text() 方法。

// 当 Scan 返回 false 时，除非已经到达文件末尾（此时 Err() 返回 nil），否则 Err() 就会返回扫描过程中出现的错误。

// 如果我使用下面命令来运行程序：

// $ go install filehandling
// $ workspacepath/bin/filehandling -fpath=/path-of-file/test.txt
// 程序会输出：

// Hello World. Welcome to file handling in Go.
// This is the second line of the file.
// We have reached the end of the file.
