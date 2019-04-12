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

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
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

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
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
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}

// 首先定义 Advertisement 类型，并在 Advertisement 类型中定义 calculate() 和 source() 方法。
// Advertisement 类型有三个字段，分别是 adName、CPC（每次点击成本）和 noOfClicks（点击次数）。
// 广告的总收益等于 CPC 和 noOfClicks 的乘积。
// 我们创建了两个广告项目，即 bannerAd 和 popupAd。incomeStream 切片包含了这两个创建的广告项目。

// 上面程序会输出：

// Income From Project 1 = $5000
// Income From Project 2 = $10000
// Income From Project 3 = $4000
// Income From Banner Ad = $1000
// Income From Popup Ad = $3750
// Net income of organisation = $23750

// 你会发现，尽管我们新增了收益流，但却完全没有修改 calculateNetIncome 函数。这就是多态带来的好处。
// 由于新的 Advertisement 同样实现了 Income 接口，所以我们能够向 incomeStreams 切片添加 Advertisement。
// calculateNetIncome 无需修改，因为它能够调用 Advertisement 类型的 calculate() 和 source() 方法。
