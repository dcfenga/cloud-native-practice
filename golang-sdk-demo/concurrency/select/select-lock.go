package main

import (
	"fmt"
	"time"
)

func main() {
	// 不加default语句时死锁, 异常：fatal error: all goroutines are asleep - deadlock!
	// 分析：由于没有 Go 协程向该信道写入数据，因此 select 语句会一直阻塞，导致死锁。
	cha := make(chan string)
	select {
	case <-cha:
	default:
		fmt.Println("default case executed1")
	}

	// 如果 select 只含有值为 nil 的信道，同样会执行默认情况
	var chb chan string
	select {
	case v := <-chb:
		fmt.Println("received value", v)
	default:

		fmt.Println("default case executed2")
	}

	// 随机选取
	// 当 select 由多个 case 准备就绪时，将会随机地选取其中之一去执行
	output1 := make(chan string)
	output2 := make(chan string)
	go server3(output1)
	go server4(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	// 输出：from server1 或者 from server2，一旦case匹配即退出

	// 空 select
	// 异常： fatal error: all goroutines are asleep - deadlock!
	// 分析： select 语句没有任何 case，因此它会一直阻塞，导致死锁
	select {}
}

func server3(ch chan string) {
	ch <- "from server1"
}

func server4(ch chan string) {
	ch <- "from server2"
}
