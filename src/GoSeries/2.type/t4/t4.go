package main

import (
	"fmt"
)

func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

// c1 和 c2 是两个复数。
// c1的实部为 5，虚部为 7。
// c2 的实部为8，虚部为 27。
// c1 和 c2 的和赋值给 cadd ，而 c1 和 c2 的乘积赋值给 cmul。
