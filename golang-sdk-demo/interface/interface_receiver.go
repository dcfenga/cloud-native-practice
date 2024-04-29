package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

/*
使用值接受者
使用值接受者声明的方法，既可以用值来调用，也能用指针调用
*/
func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

/*
使用指针接受者
*/
func (a *Address) Describe() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	// 指针接受者实现接口
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe() //输出：Sam is 25 years old

	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe() //输出：James is 32 years old

	var d2 Describer
	a := Address{"Washington", "USA"}
	/*
		异常：
		cannot use a (variable of type Address) as Describer value in assignment:
		Address does not implement Describer (method Describe has pointer receiver)
		分析：
		对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
		但接口中存储的具体值（Concrete Value）并不能取到地址，因此下面这行对于编译器无法自动获取 a 的地址，于是程序报错。
		Address 类型的指针接受者实现了接口 Describer，a 属于值类型，它并没有实现 Describer 接口，所以报错。
	*/
	//d2 = a

	d2 = &a // 这是合法的
	//Address 类型的指针实现了 Describer 接口
	d2.Describe() // 输出：State Washington Country USA
}
