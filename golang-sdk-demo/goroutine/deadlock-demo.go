package main

import (
	"fmt"
	"time"
)

func hello(pipeline chan string) {
	fmt.Println("消费者开始")
	msg := <-pipeline
	fmt.Println(msg)
	fmt.Println("消费者结束")
}

func main() {
	// 错误示例1
	// 下面代码异常：fatal error: all goroutines are asleep - deadlock!
	// 原因分析：使用 make 创建信道的时候，若不传递第二个参数，则定义的是无缓冲信道；
	// 而对于无缓冲信道，在接收者未准备好之前，发送操作是阻塞的。
	/*
		pipeline := make(chan string)
		pipeline <- "hello world"
		fmt.Println(<-pipeline)
	*/

	// 优化方法1：使接收者代码在发送者之前执行,但下面注释的代码依然发生deadlock!
	// 原因分析：发送者和接收者在同一协程中，虽然保证了接收者代码在发送者之前执行，但是由于前面接收者一直在等待数据 而处于阻塞状态；
	// 所以无法执行到后面的发送数据。还是一样造成了死锁。
	/*
		pipeline := make(chan string)
		fmt.Println(<-pipeline)
		pipeline <- "hello world"
	*/
	// 但可将接收者代码写在另一个协程里，并保证在发送者之前执行，运行正确。
	pipeline := make(chan string)
	go hello(pipeline)
	pipeline <- "hello world"
	fmt.Println("------------------")

	// 优化方法2：使用缓冲信道，而不使用无缓冲信道
	pipeline2 := make(chan string, 1)
	pipeline2 <- "hello world"
	fmt.Println(<-pipeline2)
	fmt.Println("------------------")

	// 错误示例2
	// 下面代码异常：fatal error: all goroutines are asleep - deadlock!
	// 原因分析：每个缓冲信道，都有容量，当信道里的数据量等于信道的容量后，此时再往信道里发送数据，
	// 就失造成阻塞，必须等到有人从信道中消费数据后，程序才会往下进行。
	/*
		ch1 := make(chan string, 1)
		ch1 <- "hello world"
		ch1 <- "hello china"
		fmt.Println(<-ch1)
	*/

	// 错误示例3
	// 下面代码异常：fatal error: all goroutines are asleep - deadlock!
	// 原因分析：当程序一直在等待从信道里读取数据，而此时并没有人会往信道中写入数据。此时程序就会陷入死循环，造成死锁。
	/*
		ch2 := make(chan string)
		go func() {
			pipeline <- "hello world"
			pipeline <- "hello china"
		}()

		for data := range ch2 {
			fmt.Println(data)
		}
	*/

	// 优化方法：发送完数据后，手动关闭信道，告诉 range 信道已经关闭，无需等待就行。
	ch2 := make(chan string, 1)
	go func() {
		ch2 <- "hello golang"
		ch2 <- "hello demo"
		close(ch2)
	}()

	for data := range ch2 {
		fmt.Println(data)
	}
	fmt.Println("------------------")

	time.Sleep(time.Second)
}
