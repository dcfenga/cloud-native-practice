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

type Address struct {
	city, state string
}
type Person struct {
	name    string
	age     int
	address Address
}

type Address2 struct {
	city, state string
}

type Person2 struct {
	name     string
	age      int
	Address2 //匿名字段，存在提升
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
	fmt.Println("==============\n")

	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
	fmt.Println("==============\n")

	// 提升字段（Promoted Fields）
	// 如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。
	// 这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问。
	var p2 Person2
	p2.name = "Naveen"
	p2.age = 50
	p2.Address2 = Address2{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.city)   //city is promoted field, 等价于：p2.Address2.city
	fmt.Println("State:", p2.state) //state is promoted field,等价于：p2.Address2.state
}
