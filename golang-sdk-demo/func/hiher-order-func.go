package main

import "fmt"

/*
满足下列条件之一的函数定义为高阶函数：
1.接收一个或多个函数作为参数
2.返回值是一个函数
*/

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simele2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	// 把函数作为参数，传递给其它函数
	f := func(a, b int) int {
		return a + b
	}
	simple(f) //输出：67

	// 在其它函数中返回函数
	s := simele2()        // 把 simple2函数 的返回值赋值给了 s，s 即包含了 simple 函数返回的函数
	fmt.Println(s(60, 7)) // 调用 s，并向它传递了两个 int 参数，输出：67
}
