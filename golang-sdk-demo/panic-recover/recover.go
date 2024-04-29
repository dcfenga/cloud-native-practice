package main

import (
	"fmt"
)

/*
recover 是一个内建函数，用于重新获得 panic 协程的控制。

recover 函数的标签如下所示：
func recover() interface{}

只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，
并且停止 panic 续发事件（Panicking Sequence），程序运行恢复正常。
如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。

只有在相同的 Go 协程中调用 recover 才管用。recover 不能恢复一个不同协程的 panic。
*/

func revocerName() {
	if r := recover(); r != nil {
		fmt.Println("recover from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer revocerName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("defered call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	/*输出：
	recover from  runtime error: last name cannot be nil
	returned normally from main
	defered call in main
	*/
}
