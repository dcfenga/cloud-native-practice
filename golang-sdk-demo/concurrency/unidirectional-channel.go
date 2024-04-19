package main

import (
	"fmt"
	"time"
)

/*
单向信道，可以细分为 只读信道 和 只写信道。
*/

// Sender 定义只写信道类型
type Sender = chan<- int

// Receiver 定义只读信道类型
type Receiver = <-chan int

func main() {
	// 定义只读信道
	// <-chan 表示这个信道，只能从里发出数据，对于程序来说就是只读
	var pipeline1 = make(chan int)
	type Reciver1 = <-chan int // 关键代码：定义别名类型
	var reciver1 Reciver1 = pipeline1
	fmt.Println(reciver1)

	// 定义只写信道
	// chan<- 表示这个信道，只能从外面接收数据，对于程序来说就是只写
	var pipeline2 = make(chan int)
	type Sender2 = chan<- int // 关键代码：定义别名类型
	var sender2 Sender2 = pipeline2
	fmt.Println(sender2)

	// 简写方式,不建议
	// 信道本身就是为了传输数据而存在的，如果只有接收者或者只有发送者，那信道就变成了只入不出或者只出不入了吗，没什么用。
	// 所以只读信道和只写信道，唇亡齿寒，缺一不可。
	type Sender3 chan<- int
	sender3 := make(Sender3)
	fmt.Println(sender3)
	fmt.Println("------------------")

	// 只读/只写信道测试
	var pipeline = make(chan int)

	go func() {
		var sender Sender = pipeline
		fmt.Println("准备发送数据: 100")
		sender <- 100
	}()

	go func() {
		var receiver Receiver = pipeline
		num, ok := <-receiver
		if ok {
			fmt.Println("信道还未关闭")
		} else {
			fmt.Println("信道已经关闭")
		}
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(1)
}
