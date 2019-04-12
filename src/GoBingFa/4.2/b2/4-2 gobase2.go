package main

import (
	"time"
)

func main() {
	go println("Go! Goroutine!")
	time.sleep(time.Millisecond)
}
