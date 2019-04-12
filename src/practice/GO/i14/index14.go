package main

import (
    "fmt"
	"math/rand"
)

func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4) %2 ]; interface{}(v).(type) {
	case int32 :       //可以被替换为int32或rune
		fmt.Printf("Case A.")
	case byte :
		fmt.Printf("Case B.")
	default:
		fmt.Println("Unknown!")
	}
}