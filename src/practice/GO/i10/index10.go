package main

import "fmt"

type Animal interface {
    Grow()
	Move(string) string
}
type Cat struct {
    Name     string
    Age      uint8
    Location string
}

func (cat *Cat) Grow() {
    cat.Age++
}

func (cat *Cat) Move(new string) string {
    old := cat.Location
    cat.Location = new
    return old
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}