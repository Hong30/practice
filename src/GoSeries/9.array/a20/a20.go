package main

import (
	"fmt"
)

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

// 在上述程序的第 9 行，neededCountries := countries[:len(countries)-2 创建一个去掉尾部 2 个元素的切片 countries，
// 在上述程序的 11 行，将 neededCountries 复制到 countriesCpy 同时在函数的下一行返回 countriesCpy。
// 现在 countries 数组可以被垃圾回收, 因为 neededCountries 不再被引用
