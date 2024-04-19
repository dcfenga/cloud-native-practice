package main

import "fmt"

/*
方法的接收器分：指针接收器与值接收器
值接收器和指针接收器之间的区别在于：指针接收器方法内部的改变对于调用者是可见的，然而值接收器的不行。

指针接收器使用场景1：对方法内部的接收器所做的改变应该对调用者可见时；
指针接收器使用场景2：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。
在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

在其他的所有情况，值接收器都可以被使用。
*/

type Student struct {
	name string
	age  int
}

/*
使用值接收器的方法,对 Student 结构体的字段 name 所做的改变对调用者是不可见的。
*/
func (s Student) changeName(newName string) {
	s.name = newName
}

/*
使用指针接收器的方法,对 Student 结构体的字段 age 所做的改变对调用者是可见的。
*/
func (s *Student) changeAge(newAge int) {
	s.age = newAge
}

type rectangle struct {
	length int
	width  int
}

// 值参数
func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

// 值接收器
func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

// 指针参数
func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

// 指针接收器
func main() {
	s := Student{
		name: "Mark Andrew",
		age:  16,
	}
	fmt.Printf("Student name before change: %s", s.name)
	s.changeName("Michael Andrew")
	fmt.Printf("\nStudent name after change: %s", s.name)

	fmt.Printf("\n\nStudent age before change: %d", s.age)
	(&s).changeAge(17)
	fmt.Printf("\nStudent age after change: %d", s.age)
	s.changeAge(18) //与上行代码等家，会被Go解释为 (&s).changeAge(17)
	fmt.Printf("\nStudent age after change: %d", s.age)
	fmt.Println("\n======================================")

	/* 在方法中使用值接收器与在函数中使用值参数
	   当一个函数有一个值参数，它只能接受一个值参数;
	   当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	*/
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()
	fmt.Println("\n======================================")

	p := &r //pointer to r
	/*
		异常：cannot use p (variable of type *rectangle) as rectangle value in argument to area
	*/
	// area(p)
	p.area() //成功，通过指针调用值接收器, Go语言把 p.area() 解释为 (*p).area()
	fmt.Println("\n======================================")

	/* 在方法中使用指针接收器与在函数中使用指针参数
	   和值参数相类似，函数使用指针参数只接受指针;
	   而使用指针接收器的方法可以使用值接收器和指针接收器。
	*/
	perimeter(p)  //调用 perimeter 函数时传入了一个指针参数
	p.perimeter() //通过指针接收器调用了 perimeter 方法

	/*
			异常：cannot use r (variable of type rectangle) as *rectangle value in argument to perimeter
		    分析：函数的指针参数不接受值参数
	*/
	// perimeter(r)
	r.perimeter() //使用值来调用指针接收器，Go语言把代码 r.perimeter() 解释为 (&r).perimeter()
}
