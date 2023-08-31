package main

import (
	"fmt"
)

/*
根据变量指向的值，是否是内存地址，变量分为两种：
1. 普通变量：存储数据值本身
2. 指针变量：存值的内存地址
*/
func main() {
	// &：从一个普通变量中取得内存地址
	// *：当 * 在赋值操作值的右边，是从一个指针变量中取得变量值，当*在赋值操作值的左边，是指该指针指向的变量

	// 指针创建方法1：先定义对应的变量，再通过变量取得内存地址，创建指针
	// 定义普通变量
	x := 1
	// 定义指针变量
	p := &x
	fmt.Println(p)
	fmt.Println("--------------------")

	// 指针创建方法2：先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值
	// 创建指针
	astr := new(string)
	// 给指针赋值
	*astr = "Go编程时光"
	fmt.Println(astr)
	fmt.Println(&astr)
	fmt.Println(*astr)
	fmt.Println("--------------------")

	// 指针创建方法3：先声明一个指针变量，再从其他变量取得内存地址赋值给它
	ainttemp := 1
	var binttemp *int    //声明一个指针
	binttemp = &ainttemp //初始化
	fmt.Println(ainttemp)
	fmt.Println(binttemp)
	fmt.Println(*binttemp)
	fmt.Println("--------------------")

	// &：从一个普通变量中取得内存地址
	// *：当 * 在赋值操作值的右边，是从一个指针变量中取得变量值，当*在赋值操作值的左边，是指该指针指向的变量
	y := 1 // 定义普通变量
	p = &y // 定义指针变量
	fmt.Println("普通变量存储的是：", y)
	fmt.Println("普通变量存储的是：", *p)
	fmt.Println("指针变量存储的是：", &y)
	fmt.Println("指针变量存储的是：", p)
	fmt.Printf("指针变量存储的是： %p", p)
	fmt.Println("--------------------")

	// 指针的类型
	astr2 := "hello"
	aint := 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr2 指针类型是：%T\n", &astr2)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)

	// 指针的零值
	a := 25
	var b *int //声明一个指针

	if b == nil {
		fmt.Println(b)
		b = &a //初始化：将a的内存地址给b
		fmt.Println(b)
	}

	// 指针与切片
	// 通过一个函数改变一个数组的值,方法1：将这个数组的切片做为参数传给函数
	myarr := [3]int{89, 90, 91}
	fmt.Println(myarr)
	modify1(myarr[:])
	fmt.Println(myarr)

	// 通过一个函数改变一个数组的值,方法2：将这个数组的指针做为参数传给函数
	modify2(&myarr)
	fmt.Println(myarr)
}

func modify1(sls []int) {
	sls[0] = 88
}

func modify2(arr *[3]int) {
	arr[0] = 88
}
