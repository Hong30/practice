package main

import (
	"fmt"
	"time"
)

func main() {
	name := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range name {
		go func(who string) {
			fmt.Printf("Hello, %s!\n", who)
		}(name)
	}
	time.Sleep(time.Millisecond)
}
