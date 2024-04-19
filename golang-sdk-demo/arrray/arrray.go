package main

import "fmt"

/*
数组是同一类型元素的集合。
一个数组的表示形式为 [n]T。n 表示数组中元素的数量，T 代表每个元素的类型。元素的数量 n 也是该类型的一部分。
*/
func main() {
	// 数组中的所有元素都被自动赋值为数组类型的零值
	var a [3]int //int array with length 3
	fmt.Println(a)

	// 数组的索引从 0 开始到 length - 1 结束
	var b [3]int //int array with length 3
	b[0] = 12    // array index starts at 0
	b[1] = 78
	b[2] = 50
	fmt.Println(b)

	// 简略声明 创建数组
	c := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(c)

	d := [3]int{12}
	fmt.Println(d)

	// 忽略声明数组的长度，并用 ... 代替，让编译器为你自动计算长度
	e := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(e)

	// 数组的大小是类型的一部分。因此 [5]int 和 [25]int 是不同类型
	//f := [3]int{5, 78, 8}
	//var g [5]int
	//g = f // 敷值异常：not possible since [3]int and [5]int are distinct types

	// 数组是值类型
	// Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，
	// 该变量会得到一个原始数组的一个副本。如果对新变量进行更改，则不会影响原始数组。
	m := [...]string{"USA", "China", "India", "Germany", "France"}
	n := m // a copy of a is assigned to b
	n[0] = "Singapore"
	fmt.Println("m is ", m)
	fmt.Println("n is ", n)

	// 当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	// 数组的长度
	x := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of x is", len(x))

	// for 循环遍历数组中的元素
	for i := 0; i < len(x); i++ { // looping from 0 to the length of the array
		fmt.Printf("%d th element of x is %.2f\n", i, x[i])
	}

	// range 循环遍历数组中的元素
	sum := float64(0)
	fmt.Printf("%.2f the sum is %.2f\n", sum, sum)
	for i, v := range x { //range returns both the index and value
		fmt.Printf("%d the element of x is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of x", sum)

	// 多维数组
	y := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}

	printarray(y)
	var z [3][2]string
	z[0][0] = "apple"
	z[0][1] = "samsung"
	z[1][0] = "microsoft"
	z[1][1] = "google"
	z[2][0] = "AT&T"
	z[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(z)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
