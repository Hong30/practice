package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.println("when’s Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.println("Today.")
	case today + 1:
		fmt.println("Tomorrow.")
	case today + 2:
		fmt.println("In two days.")
	default:
		fmt.println("Too far away.")
	}
}
