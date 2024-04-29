package main

import "fmt"

/*
Go 语言支持头等函数的机制,即可以把函数赋值给变量，也可以把函数作为其它函数的参数或者返回值。
*/

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {

	// 匿名函数(Anonymous Function),即没有名称的函数
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a) // 输出: func()

	// 使用()直接调用匿名函数
	func() {
		fmt.Println("hello world first class function")
	}()

	// 向匿名函数传递参数
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers") // 输出：Welcome Gophers

	/*
		自定义函数类型，类似自定义结构体
		type add func(a int, b int) int
		以上代码片段创建了一个新的函数类型 add，它接收两个整型参数，并返回一个整型。
	*/
	type add func(a int, b int) int

	var funca add = func(a int, b int) int {
		return a + b
	}
	sum := funca(5, 6)
	fmt.Println("Sum", sum) // 输出：Sum 11

	/*
		闭包（Closure）是匿名函数的一个特例。当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。
	*/
	x := 5
	func() {
		fmt.Println("x =", x) // 输出：x = 5
	}()

	// 每一个闭包都会绑定一个它自己的外围变量（Surrounding Variable）
	m := appendStr()           // 变量 m 和 n 都是闭包，它们绑定了各自的 t 值
	n := appendStr()           // 变量 m 和 n 都是闭包，它们绑定了各自的 t 值
	fmt.Println(m("World"))    // 用参数 World 调用了 m, 现在 m 中 t 值变为了 Hello World
	fmt.Println(n("Everyone")) // 用参数 Everyone 调用了 m, 由于 m 绑定了自己的变量 t，因此 m 中的 t 还是等于初始值 Hello。于是该函数调用之后，m 中的 t 变为了 Hello Everyone

	fmt.Println(m("Gopher"))
	fmt.Println(n("!"))

	/* 输出：
	Hello World
	Hello Everyone
	Hello World Gopher
	Hello Everyone !
	*/

	// 头等函数的实际用途
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	arr := []student{s1, s2}

	// 传递检查学生成绩是否为 B 的函数
	f1 := filter(arr, func(s student) bool {
		if s.grade == "B" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(f1) // 输出：[{Samuel Johnson B USA}]

	// 传递查找所有来自印度学生的函数
	f2 := filter(arr, func(s student) bool {
		if s.country == "India" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(f2) // 输出：[{Naveen Ramanathan A India}]

	// 示例：将切片中的所有整数乘以 5，并返回出结果
	// MAP函数：对集合中的每个元素进行操作的函数称为 map 函数
	z := []int{5, 6, 7, 8, 9}
	r := iMap(z, func(n int) int {
		return n * 5
	})
	fmt.Println(r) // 输出：[25 30 35 40 45]
}

// 函数 appendStr 返回了一个闭包，这个闭包绑定了变量 t
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
