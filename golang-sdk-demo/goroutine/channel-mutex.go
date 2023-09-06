package main

import (
	"fmt"
	"sync"
)

/*
当信道里的数据量已经达到设定的容量时，此时再往里发送数据会阻塞整个程序。
利用这个特性，可以用当他来当程序的锁。
*/

// 由于 x=x+1 不是原子操作
// 所以应避免多个协程对x进行操作
// 使用容量为1的信道可以达到锁的效果
func increment(ch chan bool, x *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	ch <- true
	lock.Lock()
	*x = *x + 1
	lock.Unlock()
	<-ch

	wg.Done()
}

func main() {
	// 注意要设置容量为 1 的缓冲信道
	// 如果不加锁，输出会小于1000, 加锁后输出1000
	pipeline := make(chan bool, 1)
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	wg.Add(1000)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipeline, &x, &wg, lock)
	}

	// 确保所有的协程都已完成
	// 以后会介绍一种更合适的方法（Mutex），这里暂时使用sleep
	wg.Wait()
	fmt.Println("x 的值：", x)
}
