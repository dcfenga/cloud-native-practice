package main

import (
	"dcfenga/cloud-native-practice/golang-sdk-demo/oop/employee"
	"fmt"
)

/*
Go 并不是完全面向对象的编程语言, 官网回答：
可以说是，也可以说不是。虽然 Go 有类型和方法，支持面向对象的编程风格，但却没有类型的层次结构。
Go 中的“接口”概念提供了一种不同的方法，我们认为它易于使用，也更为普遍。Go 也可以将结构体嵌套使用，
这与子类化（Subclassing）类似，但并不完全相同。此外，Go 提供的特性比 C++ 或 Java 更为通用：
子类可以由任何类型的数据来定义，甚至是内建类型（如简单的“未装箱的”整型）。
这在结构体（类）中没有受到限制。
*/
func main() {
	// 使用结构体，而非类
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining() // 输出：Sam Adolf has 10 leaves remaining
	fmt.Println("\n=============")

	// 定义一个零值的 employee 结构体变量
	var e2 employee.Employee
	e2.LeavesRemaining() // 输出：has 0 leaves remaining
	fmt.Println("\n=============")

	// 使用 New() 函数，而非构造器
	// Go 并不支持构造器。如果某类型的零值不可用，需要程序员来隐藏该类型，避免从其他包直接访问。
	e3 := employee.New("Sam", "Adolf", 30, 20)
	e3.LeavesRemaining() // 输出：Sam Adolf has 10 leaves remaining
	fmt.Println("\n=============")
}
