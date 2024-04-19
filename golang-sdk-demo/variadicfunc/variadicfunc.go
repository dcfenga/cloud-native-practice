package main

import (
	"fmt"
	"unsafe"
)

/*
可变参数函数是一种参数个数可变的函数。语法:
如果函数最后一个参数被记作 ...T，这时函数可以接受任意个 T 类型参数作为最后一个参数。

请注意只有函数的最后一个参数才允许是可变的。
*/

// 参数 nums 相当于一个整型切片,可变参数函数的工作原理是把可变参数转换为一个新的切片
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	// 没有给可变参数 nums ...int 传入任何参数，在这种情况下 nums 是一个长度和容量为 0 的 nil 切片
	find(87)
	fmt.Println("-------------")

	// 给可变参数函数传入切片
	nums := []int{89, 90, 95}
	// 异常：cannot use nums (variable of type []int) as int value in argument to find
	// 分析：func find(num int, nums ...int) 可变参数nums ...int 意味它可以接受 int 类型的可变参数
	// nums可变参数会被转换为 int 类型切片然后在传入 find 函数中；但是在这里 nums 已经是一个 int 类型切片，
	// 编译器试图在 nums 基础上再创建一个切片，即find(89, []int{nums})，nums 是一个 []int类型 而不是 int类型
	// 所以产生异常
	// find(89, nums)

	// 一种直接将切片传入可变参数函数的语法糖，即在在切片后加上 ... 后缀
	// 如果这样做，切片将直接传入函数，不再创建新的切片
	find(89, nums...)

	// 可变参数函数中的参数是切片引用传递
	welcome := []string{"hello", "world"}
	fmt.Printf("welcome addr before invoke change: %p, type: %T\n", &welcome, unsafe.Sizeof(welcome))
	change(welcome...) //以切片作为可变参数传入 change 函数
	fmt.Println("**************")
	fmt.Printf("welcome addr after invoke change: %p, type: %T\n", &welcome, unsafe.Sizeof(welcome))
	fmt.Println(welcome) //输出：[Go world]
	fmt.Println("==================================")

	fmt.Printf("welcome addr before invoke change2: %p, type: %T\n", &welcome, unsafe.Sizeof(welcome))
	change2(welcome...) //以切片作为可变参数传入 change 函数
	fmt.Println("**************")
	fmt.Printf("welcome addr after invoke change2: %p, type: %T\n", &welcome, unsafe.Sizeof(welcome))
	fmt.Println(welcome) //输出：[Go world]，原slice的内容

	/* 输出：
	welcome addr before invoke change: 0xc0000100d8, type: uintptr
	welcome addr before change s[0]: 0xc0000100f0, type: uintptr
	welcome addr before change s[0]: 0xc0000100f0, type: uintptr
	**************
	welcome addr after invoke change: 0xc0000100d8, type: uintptr
	[Go world]
	==================================
	welcome addr before invoke change2: 0xc0000100d8, type: uintptr
	welcome addr before change s[0]: 0xc000010120, type: uintptr
	welcome addr after change s[0]: 0xc000010120, type: uintptr
	welcome addr after append: 0xc000010120, type: uintptr
	[Go world playground]
	**************
	welcome addr after invoke change2: 0xc0000100d8, type: uintptr
	[Go world]
	*/
}

// s 是全新变量，内存地址是新创建的，底层数组没有重建，s指向同一个底层数组
func change(s ...string) {
	// slice不需要扩容
	fmt.Printf("welcome addr before change s[0]: %p, type: %T\n", &s, unsafe.Sizeof(s))
	s[0] = "Go"
	fmt.Printf("welcome addr before change s[0]: %p, type: %T\n", &s, unsafe.Sizeof(s))
}

// s 是全新变量，内存地址是新创建的，但是底层数据被新创建（因为扩容），s指向新建的底层数组
func change2(s ...string) {
	fmt.Printf("welcome addr before change s[0]: %p, type: %T\n", &s, unsafe.Sizeof(s))
	s[0] = "Go"
	fmt.Printf("welcome addr after change s[0]: %p, type: %T\n", &s, unsafe.Sizeof(s))
	s = append(s, "playground") //slice需要扩容，新创建一个slice，底层数组也是全新的，指针也是指向新数组的首位置
	fmt.Printf("welcome addr after append: %p, type: %T\n", &s, unsafe.Sizeof(s))
	fmt.Println(s) // 输出[Go world playground]
}
