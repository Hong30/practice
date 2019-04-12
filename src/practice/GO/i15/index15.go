package main

import (
    "fmt"
)

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
    for _, v := range []int{1, 2, 3, 4} {
        fmt.Printf("%d: %s\n", v, map1[v])
    }
}