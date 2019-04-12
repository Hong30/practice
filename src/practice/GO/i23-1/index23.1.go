package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.println("Good morning!")
	case t.Hour () < 17:
		fmt.println("Good afternoon.")
	default:
		fmt.println("Good evening.")
	}
}