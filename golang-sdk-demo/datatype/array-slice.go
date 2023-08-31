package main

import "fmt"

func main() {
	// 第一种方式，声明数组后赋值
	var arr1 [3]int
	arr1[0] = 0
	arr1[1] = 1
	arr1[2] = 2

	// 第二种方式，声明并直接初始化数组
	var arr2 [3]int = [3]int{1, 2, 3}

	// 第三种方式
	arr3 := [3]int{1, 2, 3}

	// 第四种方式
	arr4 := [...]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	// 判断数组类型
	arrA := [...]int{1, 2, 3}
	arrB := [...]int{1, 2, 3, 4}
	fmt.Printf("%d 的类型是： %T\n", arrA, arrA)
	fmt.Printf("%d 的类型是： %T\n", arrB, arrB)

	//类型字面量
	type arrAlias [3]int
	arr5 := arrAlias{7, 8, 9}
	fmt.Printf("%d 的类型是：%T", arr5, arr5)

	// 切片是对数组的一个连续片段的引用，所以切片是一个引用类型；
	// 这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集；
	// 需要注意的是，终止索引标识的项不包括在切片内（意思是这是个左闭右开的区间）
	fmt.Println("切片测试")

	// 切片构造1：对数组进行片段截取
	fmt.Printf("%d 的类型是：%T", arr5[0:2], arr5[0:2])

	// 切片构造2：从头声明赋值
	var strList []int
	strList = arr5[:]
	fmt.Printf("%d 的类型是：%T", strList, strList)
	fmt.Println("-----------------")

	// 因为切片是引用类型，默认值是nil
	var numList []int
	fmt.Println(numList)
	fmt.Println(numList == nil)
	fmt.Println("-----------------")

	// 声明一个空切片
	var numListEmpty []int
	fmt.Println(numListEmpty)
	fmt.Println("-----------------")

	// 切片构造3：使用 make函数构造,make 函数的格式：make( []Type, size, cap )
	// 关于 len 和 cap 的概念理解示例：
	// 公司名，就是变量名;
	// 公司里的所有工位，相当于已分配到的内存空间
	// 公司里的员工，相当于元素
	// cap 代表你这个公司最多可以容纳多少员工
	// len 代表你这个公司当前有多少个员工，所以cap>=len
	a := make([]int, 2)
	b := make([]int, 2, 6)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))

	// 切片中添加元素
	var myslice = []int{1} // 追加一个元素
	fmt.Println(myslice)
	fmt.Println("-----------------")

	myslice = append(myslice, 2)
	fmt.Println(myslice)
	fmt.Println("-----------------")

	// 追加多个元素
	myslice = append(myslice, 3, 4)
	fmt.Println(myslice)
	fmt.Println("-----------------")

	// 追加一个切片, ... 表示解包，不能省略
	myslice = append(myslice, []int{7, 8}...)
	fmt.Println(myslice)
	fmt.Println("-----------------")

	// 在第一个位置插入元素
	myslice = append([]int{0}, myslice...)
	fmt.Println(myslice)
	fmt.Println("-----------------")

	// 在中间插入一个切片(两个元素)
	myslice = append(myslice[:5], append([]int{5, 6}, myslice[5:]...)...)
	fmt.Println(myslice)
	fmt.Println("-----------------")

	// 切片自动扩容
	var number = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mynumber := number[4:6:8]
	fmt.Printf("mynumber为 %d, 其长度为: %d, 其容量为: %d\n", mynumber, len(mynumber), cap(mynumber))

	mynumber = mynumber[:cap(mynumber)]
	// 最大只能取到第4个元素，取第5个元素下标越界
	fmt.Printf("mynumber的第四个元素为: %d", mynumber[3])
}
