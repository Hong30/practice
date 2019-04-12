package main

import (
	"fmt"
	"time"
)

func main() {
	name := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range name {
		go func() {
			fmt.Printf("Hello, %s!\n", name)
		}()
	}
	time.Sleep(time.Millisecond)
}
