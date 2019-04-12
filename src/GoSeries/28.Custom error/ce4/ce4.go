package main

import "fmt"

type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

// 从 Error() string 方法中返回了关于错误的描述。
// 当 length 小于零时，lengthNegative() bool 方法返回 true，而当 width 小于零时，widthNegative() bool 方法返回 true。
// 这两个方法都提供了关于错误的更多信息，在这里，它提示我们计算面积失败的原因（长度为负数或者宽度为负数）。
// 于是我们就有了两个错误类型结构体的方法，来提供更多的错误信息。
// rectArea 函数检查了长或宽是否小于零，如果小于零，rectArea 会返回一个错误信息，
// 否则 rectArea 会返回矩形的面积和一个值为 nil 的错误。
// 在 main 程序中，我们检查了错误是否为 nil（第 4 行）。
// 如果错误值不是 nil，我们会在下一行断言 *areaError 类型。
// 然后，我们使用 lengthNegative() 和 widthNegative() 方法，检查错误的原因是长度小于零还是宽度小于零。
// 这样我们就使用了错误结构体类型的方法，来提供更多的错误信息。
// 如果没有错误发生，就会打印矩形的面积。

// 该程序会打印输出：

// error: length -5.00 is less than zero
// error: width -9.00 is less than zero
