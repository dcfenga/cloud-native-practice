package main

import (
	"fmt"
	"time"
)

/*
缓冲信道:
只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。
通过向 make 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。
ch := make(chan type, capacity)

无缓冲信道的发送和接收过程是阻塞的。要让一个信道有缓冲，上面语法中的 capacity 应该大于 0。
无缓冲信道的容量默认为 0, 该参数可省略。
*/
func main() {
	ch := make(chan string, 2) //信道的容量为 2
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 示例：一个并发的 Go 协程来向信道写入数据，而 Go 主协程负责读取数据
	cha := make(chan int, 2)
	go write(cha)
	time.Sleep(2 * time.Second)
	for v := range cha {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
	/*输出:
	successfully wrote 0 to ch
	successfully wrote 1 to ch
	read value 0 from ch
	successfully wrote 2 to ch
	read value 1 from ch
	successfully wrote 3 to ch
	read value 2 from ch
	successfully wrote 4 to ch
	read value 3 from ch
	read value 4 from ch
	*/

	// 死锁
	chb := make(chan string, 2)
	chb <- "naveen"
	chb <- "paul"

	// 异常：fatal error: all goroutines are asleep - deadlock!
	// 分析：向容量为 2 的缓冲信道写入 3 个字符串。当在程序控制到达第 3 次写入时，由于它超出了信道的容量，因此这次写入发生了阻塞。
	// chb <- "steve"
	fmt.Println(<-chb)
	fmt.Println(<-chb)

	// 长度 vs 容量
	// 缓冲信道的容量是指信道可以存储的值的数量。使用 make 函数创建缓冲信道的时候会指定容量大小。
	// 缓冲信道的长度是指信道中当前排队的元素个数。
	chc := make(chan string, 3)
	chc <- "naveen"
	chc <- "paul"
	fmt.Println("capacity is", cap(chc))
	fmt.Println("length is", len(chc))
	fmt.Println("read value", <-chc)
	fmt.Println("new length is", len(chc))
	/*输出：
	capacity is 3
	length is 2
	read value naveen
	new length is 1
	*/
}

func write(cha chan int) {
	for i := 0; i < 5; i++ {
		cha <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(cha)
}
