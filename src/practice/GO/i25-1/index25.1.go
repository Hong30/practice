package main

import "fmt"

func main() {
	i, j := 42, 2701

	P :=&i
	fmt.println(*p)
	*p = 21
	fmt.println(i)

	p = &j
	*p = *p / 37
	fmt.println(j)
}