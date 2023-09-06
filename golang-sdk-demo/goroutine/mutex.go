package main

import (
	"fmt"
	"sync"
	"time"
)

/*
在 Go 语言中，信道的地位非常高，它是 first class 级别的，面对并发问题，
我们始终应该优先考虑使用信道，如果通过信道解决不了的，不得不使用共享内存来实现并发编程的，
那 Golang 中的锁机制,就是绕不过的知识点了。

sync 包提供的锁类型
(1)一个叫 Mutex， 利用它可以实现互斥锁; Mutex，全称 mutual exclusion）是为了来保护一个资源不会因为并发操作而引起冲突导致数据不准确。
(2)一个叫 RWMutex，利用它可以实现读写锁。

Mutext 锁使用注意点：
(1)同一协程里，不要在尚未解锁时再次使加锁
(2)同一协程里，不要对已解锁的锁再次解锁
(3)加了锁后，别忘了解锁，必要时使用 defer 语句
*/
func main() {
	// 互斥锁定义方式1
	var l1 *sync.Mutex
	l1 = new(sync.Mutex)
	fmt.Println(l1)
	fmt.Println("-------------")

	// 互斥锁定义方式2
	l2 := &sync.Mutex{}
	fmt.Println(l2)
	fmt.Println("-------------")

	// 互斥锁示例1
	var wg1 sync.WaitGroup
	count1 := 0
	wg1.Add(3)
	go add1(&count1, &wg1)
	go add1(&count1, &wg1)
	go add1(&count1, &wg1)

	wg1.Wait()
	// 输出结果有可能不是3000，是随机的，原因在于：这三个协程在执行时，先读取 count 再更新 count 的值，
	// 而这个过程并不具备原子性，所以导致了数据的不准确。
	fmt.Println("count的值为：", count1)
	fmt.Println("-------------")

	// 互斥锁示例2，通过锁机制解决示例1的缺陷
	var wg2 sync.WaitGroup
	lock := &sync.Mutex{}
	count2 := 0
	wg2.Add(3)
	go add2(&count2, &wg2, lock)
	go add2(&count2, &wg2, lock)
	go add2(&count2, &wg2, lock)

	wg2.Wait()
	fmt.Println("count的值为：", count2)
	fmt.Println("-------------")

	// 读写锁：RWMutex, 将程序对资源的访问分为读操作和写操作
	// 为了保证数据的安全，它规定了当有人还在读取数据（即读锁占用）时，不允计有人更新这个数据（即写锁会阻塞）
	// 为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像 Mutex 那样只允许有一个人（线程）读取同一个数据。
	// RWMutex 里提供了两种锁，每种锁分别对应两个方法，为了避免死锁，两个方法应成对出现，必要时请使用 defer。
	// 读锁：调用 RLock 方法开启锁，调用 RUnlock 释放锁
	// 写锁：调用 Lock 方法开启锁，调用 Unlock 释放锁（和 Mutex类似）

	// 读写锁定义方式1
	var l3 *sync.RWMutex
	l3 = new(sync.RWMutex)
	fmt.Println(l3)
	fmt.Println("-------------")

	// 读写锁定义方式2
	l4 := &sync.RWMutex{}
	fmt.Println(l4)
	fmt.Println("-------------")

	// 读写锁示例
	l5 := &sync.RWMutex{}
	l5.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			l5.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			l5.RUnlock()
		}(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	l5.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	l5.Lock()
	fmt.Println("程序退出...")
	l5.Unlock()
}

func add1(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count = *count + 1
	}
	wg.Done()
}

func add2(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}
