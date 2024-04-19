package main

import "fmt"

/*
指针是一种存储变量内存地址（Memory Address）的变量
*/
func main() {

	//指针的声明: 指针变量的类型为 *T，该指针指向一个 T 类型的变量
	b := 255
	var a *int = &b
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// 指针的零值
	c := 25
	var d *int
	if d == nil {
		fmt.Println("c is", c)
		d = &c
		fmt.Println("d after initialization is", d)
	}

	// 指针的解引用, 可以获取指针所指向的变量的值,将 a 解引用的语法是 *a
	e := 255
	f := &e
	fmt.Println("address of f is", f)
	fmt.Println("value of f is", *f)

	// 用指针修改变量的值
	g := 255
	h := &g
	fmt.Println("address of g is", h)
	fmt.Println("value of g is", *h)
	*h++
	fmt.Println("new value of g is", g)

	//向函数传递指针参数
	m := 58
	fmt.Println("value of m before function call is", m)
	n := &m
	change(n)
	fmt.Println("value of m after function call is", m)
	fmt.Println("-----------")

	// 不建议的做法(不是 Go 语言惯用的实现方式s)：向函数传递数组的指针
	x := [3]int{89, 90, 91}
	modify(&x)
	fmt.Println(x)
	fmt.Println("-----------")

	// 建议的做法：向函数传递数组的切片
	y := [3]int{89, 90, 91}
	modify2(y[:])
	fmt.Println(y)
	fmt.Println("-----------")

	// Go 不支持指针运算
	p := &y
	fmt.Println(*p)

	// 异常：invalid operation: p++ (non-numeric type *int)
	// p++
}

func modify2(arr []int) {
	arr[0] = 90
}

func modify(arr *[3]int) {
	// 建议写法
	(*arr)[0] = 90
	// 简化写法，a[x] 是 (*a)[x] 的简写形式
	arr[1] = 100
}

func change(n *int) {
	*n = 55
}
