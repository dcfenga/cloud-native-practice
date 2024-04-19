package main

import (
	"fmt"
	"time"
)

/*
select 语句用于在多个发送/接收信道操作中进行选择。select 语句会一直阻塞，直到发送/接收操作准备就绪。
如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。
该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作。
*/
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s1 := <-output2:
		fmt.Println(s1)
	}
	// 输出：from server2，case一旦匹配，程序即终止

	// 默认情况
	// 在没有 case 准备就绪时，可以执行 select 语句中的默认情况（Default Case）。
	// 这通常用于防止 select 语句一直阻塞。
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

	/*输出：
	no value received
	no value received
	no value received
	no value received
	no value received
	no value received
	no value received
	no value received
	no value received
	no value received
	received value:  process successful
	*/
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
