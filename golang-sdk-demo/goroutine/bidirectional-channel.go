package main

import (
	"fmt"
	"time"
)

func main() {
	// 默认定义的信道都是双向通道，可发送数据，也可以接收数据。
	pipeline := make(chan int)
	defer close(pipeline)

	go func() {
		fmt.Println("准备发送数据: 100")
		pipeline <- 100
	}()

	go func() {
		time.Sleep(time.Second)
		num := <-pipeline
		fmt.Printf("接收到的数据是: %d", num)
	}()

	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(2 * time.Second)
}
