package main

import "fmt"

/*
函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。

函数的声明:

		func 函数名(形式参数列表)(返回值列表){
		    函数体
		}

	    形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由参数调用者提供；
	    返回值列表描述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。
*/
func sum(a int, b int) int {
	return a + b
}

// 多个类型一致的参数,参数个数可变,使用...来接收多个参数
func sumx(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func sumy(args ...int) int {
	var res int
	for _, v := range args {
		res += v
	}
	return res
}

// Sumy 使用 ...  来接收多个参数，除此之外，它还有一个用法，就是用来解序列，
// 将函数的可变参数（一个切片）一个一个取出来，传递给另一个可变参数的函数，而不是传递可变参数变量本身。
func Sumy(args ...int) int {
	// 利用 ... 来解序列
	res := sumy(args...)
	return res
}

// MyPrintf 多个类型不一致的参数,参数个数可变,使用...来接收多个参数
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 返回值1： Go 支持一个函数返回多个值
func double1(a int) (int, int) {
	b := a * 3
	return a, b
}

// 返回值2： Go支持返回带有变量名的值
func double2(a int) (b int) {
	// 不能使用 := ,因为在返回值哪里已经声明了为int
	b = a * 3
	// 不需要指明写回哪个变量,在返回值类型那里已经指定了
	return
}

/*
匿名函数，就是没有名字的函数，它只有函数逻辑体，而没有函数名。

定义的格式如下:

	func(参数列表)(返回参数列表){
	    函数体
	}
*/
func visit(list []int, f func(int)) {
	for _, v := range list {
		// 执行回调函数
		f(v)
	}
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*
方法和函数有什么区别？
方法，是一种特殊的函数。当你一个函数和对象/结构体进行绑定的时候，我们就称这个函数是一个方法。
*/
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println("-----------------")

	fmt.Println(sumx(1, 2, 3))
	fmt.Println("-----------------")

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
	fmt.Println("-----------------")

	fmt.Println(Sumy(1, 2, 3))
	fmt.Println("-----------------")

	// 接收参数用逗号分隔
	a, b := double1(2)
	fmt.Println(a, b)
	fmt.Println("-----------------")

	fmt.Println(double2(2))
	fmt.Println("-----------------")

	// 匿名函数
	func(data int) {
		fmt.Println("hello", data)
	}(100)
	fmt.Println("-----------------")

	// 使用匿名函数直接做为参数
	visit([]int{1, 2, 3, 4, 5}, func(v int) {
		fmt.Println(v)
	})
	fmt.Println("-----------------")

	// 空白符，可以用作表示任何类型的任何值
	area, _ := rectProps(10.8, 5.6) // 返回值周长被丢弃
	fmt.Printf("Area %f ", area)
}
