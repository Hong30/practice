// package main

// import (
//     "fmt"
// )

// func main() {
//     byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
//     str := string(byteSlice)
//     fmt.Println(str)
// }

package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}
