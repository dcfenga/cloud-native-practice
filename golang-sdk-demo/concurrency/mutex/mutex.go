package main

import (
	"fmt"
	"sync"
)

/*
临界区（Critical Section）概念:
当程序并发地运行时，多个 Go 协程不应该同时访问那些修改共享资源的代码。这些修改共享资源的代码称为临界区。

Mutex：
Mutex 用于提供一种加锁机制（Locking Mechanism），可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。

Mutex 可以在 sync 包内找到。Mutex定义了两个方法：Lock 和 Unlock。所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。

mutex.Lock()
x = x + 1
mutex.Unlock()

如果有一个 Go 协程已经持有了锁（Lock），当其他协程试图获得该锁时，这些协程会被阻塞，直到 Mutex 解除锁定为止。
*/

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

var y = 0

func increment2(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

var z = 0

// 容量为1的缓冲信道用于保证只有一个协程访问增加 x 的临界区
func increment3(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}

func main() {
	// 含有竞态条件的程序
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x:", x)
	// 输出：由于存在竞态条件，每一次输出都不同

	// 使用 Mutex处理含有竞态条件的程序
	var w2 sync.WaitGroup
	var m2 sync.Mutex
	for i := 0; i < 1000; i++ {
		w2.Add(1)
		// 需传递 Mutex 的地址，如果传递的是 Mutex 的值，而非地址，那么每个协程都会得到 Mutex 的一份拷贝，竞态条件还是会发生。
		go increment2(&w2, &m2)
	}
	w2.Wait()
	fmt.Println("final value of y:", y) // 输出：final value of y: 1000

	// 使用信道处理竞态条件
	var w3 sync.WaitGroup
	chc := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w3.Add(1)
		go increment3(&w3, chc)
	}
	w3.Wait()
	fmt.Println("final value of z:", z) // 输出：final value of z: 1000

	// Mutex vs 信道
	// 当 Go 协程需要与其他协程通信时，可以使用信道; 而当只允许一个协程访问临界区时，可以使用 Mutex。
	// 选择针对问题的工具，而别让问题去将就工具。
}
