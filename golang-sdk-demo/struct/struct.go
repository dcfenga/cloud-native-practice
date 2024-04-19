package main

import "fmt"

/*
结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
1. 声明结构体
type 结构体名 struct {
    属性名   属性类型// 结构体声明：命名的结构体（Named Structure）
	属性名   属性类型type Employee struct {
	...
}
2. 定义方法
使用组合函数的方式来定义结构体方法

3. 方法可见性
(1)当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用;
(2)当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
*/

// 结构体声明：命名的结构体（Named Structure）
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// 结构体声明: 简化写法，命名的结构体（Named Structure）
type Employee2 struct {
	firstName, lastName string
	age, salary         int
}

// 结构体声明: 匿名结构体，即声明结构体时不用声明一个新类型
var employee struct {
	firstName, lastName string
	age                 int
}

// 结构体声明: 匿名字段, 创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。
// 虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型; Person 结构体有两个名为 string 和 int 的字段。
type PersonS struct {
	string
	int
}

func main() {
	//通过指定每个字段名的值，字段名的顺序不一定要与声明结构体类型时的顺序相同
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	//省略了字段名，在这种情况下，就需要保证字段名的顺序与声明结构体时的顺序相同
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// 创建匿名结构体
	// 之所以称这种结构体是匿名的，是因为它只是创建一个新的结构体变量 emp3，而没有定义任何结构体类型。
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	// 结构体的零值（Zero Value）
	var emp4 Employee               //zero valued structure
	fmt.Println("Employee 4", emp4) //输出：Employee 4 {  0 0}
	// 因为emp4没有初始化任何值，因此 firstName 和 lastName 赋值为 string 的零值（""），而 age 和 salary 赋值为 int 的零值（0）

	// 可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值。
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5) //输出：Employee 5 {John Paul 0 0}

	// 匿名字段
	p := PersonS{"Naveen", 50}
	fmt.Println(p)
	fmt.Println("=================")

	// Person 结构体有两个名为 string 和 int 的字段
	var p1 PersonS
	p1.string = "FDC"
	p1.int = 35
	fmt.Println(p)
}
