package main

import "fmt"

/*
Go 语言本身并不支持继承,但可以使用组合的方法，实现类似继承的效果；
在 Go 语言中，把一个结构体嵌入到另一个结构体的方法，称之为组合。
*/

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
}

func main() {
	myCom := company{
		companyName: "DCITS",
		companyAddr: "神州数码科技园",
	}

	staffInfo := staff{
		name:     "小明",
		age:      30,
		gender:   "男",
		position: "云计算开发工程师",
		company:  myCom,
	}

	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}
