package main

import (
	"fmt"
	"sync"
	"time"
)

/*
“不要通过共享内存来通信，要通过通信来共享内存”
*/

func main() {
	// 保证子goroutine执行方法1： time.Sleep 方式
	// 信道可以实现多个协程间的通信，那么我们只要定义一个信道，在任务完成后，往信道中写入true，
	// 然后在主协程中获取到true，就认为子协程已经执行完毕。
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	time.Sleep(time.Second)
	fmt.Println("------------------")

	// 保证子goroutine执行方法2： 使用 WaitGroup
	// WaitGroup 实例化后有下面几个方法：
	// Add：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量
	// Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
	// Wait：阻塞当前协程，直到实例里的计数器归零。
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
}

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}
