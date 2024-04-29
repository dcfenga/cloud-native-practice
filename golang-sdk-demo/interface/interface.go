package main

import "fmt"

/*
接口定义：面向对象的领域里，接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。

在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。
*/

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString 实现了 VowelsFinder 接口
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay * p.pf
}

type Contract struct {
	empId    int
	basicpay int
}

// salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

// 接收空接口作为参数，因此，可以给这个函数传递任何类型
func describeBlank(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType2(i interface{}) {
	switch v := i.(type) {
	// v 与接口类型 Describer 进行了比较。p 实现了 Describer，因此满足了该 case 语句
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func assert2(x interface{}) {
	v, ok := x.(int)
	fmt.Println(v, ok)
}

func assert(x interface{}) {
	s := x.(int) //get the underlying int value from x
	fmt.Println(s)
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(employees []SalaryCalculator) {
	expense := 0
	for _, v := range employees {
		// 通过调用不同类型对应的 CalculateSalary 方法，totalExpense 可以计算得到支出
		// 优点：totalExpense 可以扩展新的员工类型，而不需要修改任何代码
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

func main() {
	// 接口的声明与实现
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	// 接口价值示例
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	/*
		接口的内部表示
		可以把接口看作内部的一个元组 (type, value)。type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。
	*/
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t) // t 的具体类型为 MyFloat
	t.Tester()  // t 的值为 89.7

	/*
		空接口:没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口。
	*/
	s := "Hello World"
	describeBlank(s)
	i := 55
	describeBlank(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describeBlank(strt)
	/*输出
	Type = string, value = Hello World
	Type = int, value = 55
	Type = struct { name string }, value = {Naveen R}
	*/

	/*
		类型断言：
		类型断言用于提取接口的底层值（Underlying Value）。
		在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
	*/
	var x interface{} = 56
	// 异常：panic: interface conversion: interface {} is string, not int
	assert(x)

	// var y interface{} = "Steven Paul"
	// 异常：panic: interface conversion: interface {} is string, not int
	// assert(y)

	/*
		使用下面语法避免上述异常
		v, ok := i.(T)
		如果 i 的具体类型是 T，那么 v 赋值为 i 的底层值，而 ok 赋值为 true。
		如果 i 的具体类型不是 T，那么 ok 赋值为 false，v 赋值为 T 类型的零值，此时程序不会报错。
	*/
	var x2 interface{} = 56
	// 异常：panic: interface conversion: interface {} is string, not int
	assert2(x2) //输出： 56  true

	var y2 interface{} = "Steven Paul"
	assert2(y2) //输出： 0  false

	/*
		类型选择（Type Switch）
		类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。
		唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。

		类型选择的语法类似于类型断言。类型断言的语法是 i.(T)，而对于类型选择，类型 T 由关键字 type 代替。
	*/
	findType("Naveen")
	findType(77)
	findType(89.98)

	// 将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
	findType2("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p) // 输出：Naveen R is 25 years old
}
