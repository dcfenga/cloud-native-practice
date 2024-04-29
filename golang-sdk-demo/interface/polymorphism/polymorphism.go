package main

import "fmt"

/*
使用接口实现多态

一个类型如果定义了接口的所有方法，那它就隐式地实现了该接口。

所有实现了接口的类型，都可以把它的值保存在一个接口类型的变量中。在 Go 中，我们使用接口的这种特性来实现多态。
*/

// 示例：计算一个组织机构的净收益，该组织所获得的收入来源于两个项目之和：fixed billing 和 time and material
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

/*
FixedBilling 和 TimeAndMaterial 两个结构体都定义了 Income 接口的两个方法：calculate() 和 source()，
因此这两个结构体都实现了 Income 接口。
*/

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

/*
calcuateNetIncome函数接收一个 Income 接口类型的切片作为参数。
该函数会遍历这个接口切片，并依个调用 calculate() 方法，计算出总收入。
该函数同样也会通过调用 source() 显示收入来源。根据 Income 接口的具体类型，
程序会调用不同的 calculate() 和 source() 方法。于是，我们在 calculateNetIncome 函数中就实现了多态。
*/
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

	/*输出：
	Income From Project 1 = $5000
	Income From Project 2 = $10000
	Income From Project 3 = $4000
	Net income of organisation = $19000
	*/

	// 新增收益流：Advertisement
	project11 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project22 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project33 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams2 := []Income{project11, project22, project33, bannerAd, popupAd}
	calculateNetIncome(incomeStreams2)

	/*输出：
	Income From Project 1 = $5000
	Income From Project 2 = $10000
	Income From Project 3 = $4000
	Net income of organisation = $19000Income From Project 1 = $5000
	Income From Project 2 = $10000
	Income From Project 3 = $4000
	Income From Banner Ad = $1000
	Income From Popup Ad = $3750
	Net income of organisation = $23750
	*/
}
