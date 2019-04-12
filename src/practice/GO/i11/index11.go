package main

import "fmt"

type MyInt struct {
    n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
}

func main() {
	mi := MyInt{}
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)
}