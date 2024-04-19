package main

import (
	"fmt"
	"time"
)

/*
信道：
信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，
通过使用信道，数据也可以从一端发送，在另一端接收。
*/
func main() {

	// 信道的声明
	// 所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
	// chan T 表示 T 类型的信道
	// 信道的零值为 nil。信道的零值没有什么用，应该像对 map 和切片所做的那样，用 make 来定义信道。
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("type if a is %T\n", a)
		fmt.Println("=====================")
	}

	// 简短声明
	// a: = make(chan int)

	/* 通过信道进行发送和接收,语法如下：
	   data := <- a // 读取信道 a
	   a <- data    // 写入信道 a
	   信道旁的箭头方向指定了是发送数据还是接收数据
	   判断规则：<-在信道的左侧为读，<- 在信道的右侧为写
	*/

	/*
	   发送与接收默认是阻塞的：
	   把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
	   与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。

	   信道的这种特性能够帮助 Go 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。
	*/

	// 主协程休眠等到hello协程结束
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function1")
	fmt.Println("=====================")

	// 使用信道来重写上面代码
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello2(done)
	// <-done 这行代码通过协程 done 接收数据，但并没有使用数据或者把数据存储到变量中,合法的用法。
	<-done
	fmt.Println("Main received data")

	// 信道示例：计算一个数中每一位的平方和与立方和，然后把平方和与立方和相加并打印出来
	number := 589
	sqrch := make(chan int)
	cubch := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubch)
	squares, cubes := <-sqrch, <-cubch
	fmt.Println("Final output", squares+cubes)

	// 死锁
	// 当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。如果没有的话，程序就会在运行时触发 panic，形成死锁。
	// ch := make(chan int)
	// ch <- 5 // 异常：fatal error: all goroutines are asleep - deadlock!

	// 单向信道，即只能发送或者接收数据
	// chan<- int 定义了唯送信道，因为箭头指向了 chan
	sendch := make(chan<- int)
	go sendData(sendch)

	// 异常：invalid operation: cannot receive from send-only channel sendch (variable of type chan<- int)
	// 分析：从唯送信道接收数据编译器报错
	// fmt.Println(<-sendch)

	// 信道转换（Channel Conversion）：把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行。
	cha := make(chan int)
	go sendData(cha)
	fmt.Println(<-cha)
}

// 把 ch 转换为一个唯送信道, 在 sendData 协程里是一个唯送信道，而在 Go 主协程里是一个双向信道。
func sendData(ch chan<- int) {
	ch <- 10
}

func calcSquares(number int, sqrch chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	sqrch <- sum //结果发送给信道
}

func calcCubes(number int, cubch chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubch <- sum //结果发送给信道
}

func hello2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func hello() {
	fmt.Println("hello world goroutine")
}
