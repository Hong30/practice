package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}

// 定义了接口 Interface，它包含了两个方法：calculate() 计算并返回项目的收入，而 source() 返回项目名称。
// 项目 FixedBillin 有两个字段：projectName 表示项目名称，而 biddedAmount 表示组织向该项目投标的金额。
// TimeAndMaterial 结构体用于表示项目 Time and Material。
// 结构体 TimeAndMaterial 拥有三个字段名：projectName、noOfHours 和 hourlyRate。
// 在项目 FixedBilling 里面，收入就是项目的投标金额。因此我们返回 FixedBilling 类型的 calculate() 方法。
// 而在项目 TimeAndMaterial 里面，收入等于 noOfHours 和 hourlyRate 的乘积，作为 TimeAndMaterial 类型的 calculate() 方法的返回值。
// 我们还通过 source() 方法返回了表示收入来源的项目名称。
// 由于 FixedBilling 和 TimeAndMaterial 两个结构体都定义了 Income 接口的两个方法：calculate() 和 source()，因此这两个结构体都实现了 Income 接口。
// calculateNetIncome函数接收一个 Income 接口类型的切片作为参数。
// 该函数会遍历这个接口切片，并依个调用 calculate() 方法，计算出总收入。
// 该函数同样也会通过调用 source() 显示收入来源。根据 Income 接口的具体类型，程序会调用不同的 calculate() 和 source() 方法。
// 于是，我们在 calculateNetIncome 函数中就实现了多态。
// 如果在该组织以后增加了新的收入来源，calculateNetIncome 无需修改一行代码，就可以正确地计算总收入了。
// main 函数中，我们创建了三个项目，有两个是 FixedBilling 类型，一个是 TimeAndMaterial 类型。
// 接着我们创建了一个 Income 类型的切片，存放了这三个项目。由于这三个项目都实现了 Interface 接口，因此可以把这三个项目放入 Income 切片。
// 最后我们将该切片作为参数，调用了 calculateNetIncome 函数，显示了项目不同的收益和收入来源。

// 该程序会输出：

// Income From Project 1 = $5000
// Income From Project 2 = $10000
// Income From Project 3 = $4000
// Net income of organisation = $19000
