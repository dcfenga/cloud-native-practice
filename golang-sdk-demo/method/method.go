package main

import (
	"fmt"
	"math"
)

/*
方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。

创建一个方法语法:
func (t Type) methodName(parameter list) {
}

函数与方法的区别：
1. Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径；
2. 相同名字的方法可以定义在不同的类型上，而相同名字的函数却是不被允许的。
   假设我们有一个 Square 和 Circle 结构体；可以在 Square 和 Circle 上分别定义一个 Area 方法。见下面的程序。
*/

type Employee struct {
	name     string
	salary   int
	currency string
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

type person struct {
	firstName string
	lastName  string
	address
}

type address struct {
	city  string
	state string
}

// 匿名字段的方法
func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// 异常：cannot define new methods on non-local type int
// 分析：为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中；
// 当前add 方法的定义和 int 类型的定义不在同一个包中，解决方法是为内置类型 int 创建一个类型别名，
// 然后创建一个以该类型别名为接收器的方法。
//func (a int) add(b int) {
//	fmt.Printf("a+b=%d", a+b)
//}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法

	displaySalary(emp1) // 调用displaySalary() 方法,  Employee 实例以参数传入

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	// 属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样
	p.fullAddress()         //address 结构体的 fullAddress 方法
	p.address.fullAddress() //不建议写法
	fmt.Println("==============")

	// 在非结构体上的方法
	// 为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
