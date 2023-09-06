package main

import (
	"fmt"
	"time"
)

/*
一个 goroutine 本身就是一个函数，当你直接调用时，它就是一个普通函数，
如果你在调用前加一个关键字 go ，那你就开启了一个 goroutine。

函数main 的地位相当于主线程，当 main 函数执行完成后，这个线程也就终结了，
其下的运行着的所有协程也不管代码是不是还在跑，也得乖乖退出。
*/

func mygoroutine() {
	fmt.Println("hello, go")
}

func main() {
	go mygoroutine()
	fmt.Println("hello, world")

	// 需加上这行代码，因为协程的创建需要时间，当 hello, world打印后，协程还没来得及并执行，主进程就退出了
	time.Sleep(time.Second)
}
