package main

import "fmt"

/*
变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。
一般变量分为两种 普通变量 和 指针变量；普通变量，存放的是数据本身，而指针变量存放的是数据的地址。
*/
func main() {
	// 第一种：声明单个变量,一行一个变量，静态语言最基本常用的方式
	// 可用于常量，替换关键字 var -> const 即可
	// 声明变量未赋值时，Go 会自动地将其初始化，赋值该变量类型的零值（Zero Value）。
	var age1 int
	var name1 string = "Go编程时光发布在Go语言中文网1"
	fmt.Println(name1, age1)

	// 第二种： 变体，通过类型推导
	var age2 int
	age2 = 30
	var name2 = "Go编程时光发布在Go语言中文网2"
	fmt.Println(name2, age2)

	// 第三种：多个变量一起声明，声明组
	// 可用于常量，替换关键字 var -> const 即可
	var (
		name   = "张三"
		age    = 30
		gender = "男"
	)
	fmt.Println(name, age, gender)

	// 第三种：短声明，只能在函数内
	name3 := "Go编程时光发布在Go语言中文网3"
	fmt.Println(name3)

	// 第四种：短声明的变体，一行声明和初始化多个变量
	name4, age4 := "Go编程时光发布在Go语言中文网4", 20
	fmt.Println(name4, age4)

	// 简短声明要求 := 操作符左边的所有变量都有初始值。
	// 下面代码异常：assignment mismatch: 2 variables but 1 value
	//name5, age5 := "fdc"
	//fmt.Println(name5, age5)

	// 简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的。
	x, y := 20, 30
	fmt.Println(x, y)
	y, z := 40, 50
	fmt.Println(y, z)
	// 下面代码异常：no new variables on left side of :=
	//y, z := 60, 70
	//fmt.Println(y, z)

	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)

	// 第五种：通过 new 创建指针变量
	var num int = 28
	var ptr = &num // & 后面接变量名，表示取出该变量的内存地址
	fmt.Println("num: ", num)
	fmt.Println("ptr: ", ptr)

	// 使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，
	// 然后返回变量地址，返回的指针类型为*Type;
	// 下面两种写法等价：
	//
	// 使用 new
	// func newInt() *int {
	//    return new(int)
	// }
	//
	// 使用传统的方式
	// func newInt() *int {
	//    var dummy int
	//    return &dummy
	// }
	ptr2 := new(int)
	fmt.Println("ptr address: ", ptr2)
	fmt.Println("ptr value: ", *ptr2) // * 后面接指针变量，表示从内存地址中取出值

	// 第六种：make 函数创建 slice、map 或 chan 类型变量
	// chan 只能用 make
	// 简化：slice := []int{0, 0}
	// 简化：m := map[string]int{}
	var slice = make([]int, 4)
	var m = make(map[string]int)
	var c = make(chan int)
	fmt.Println("slice", slice)
	fmt.Println("map", m)
	fmt.Println("chan", c)

	// 第七种：匿名变量，也称作占位符，或者空白标识符，用下划线表示，优点有三：
	// 1.不分配内存，不占用内存空间
	// 2.不需要你为命名无用的变量名而纠结
	// 3.多次声明不会有任何问题

	// Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值。
	// 下面的程序会抛出错误 cannot use "fdc" (untyped string constant) as int value in assignment.
	//w := 29
	//w = "fdc"
	//fmt.Println(w)w
}
