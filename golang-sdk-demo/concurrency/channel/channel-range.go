package main

import (
	"fmt"
)

/*
数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
v, ok := <- ch
如果成功接收信道所发送的数据，那么 ok 等于 true。而如果 ok 等于 false，说明我们试图读取一个关闭的通道。
从关闭的信道读取到的值会是该信道类型的零值。例如，当信道是一个 int 类型的信道时，那么从关闭的信道读取的值将会是 0。
*/
func main() {

	chb := make(chan int)
	go producer(chb)
	for {
		v, ok := <-chb
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	// for range 循环重写上面的代码
	// for range 循环用于在一个信道关闭之前，从信道接收数据
	chc := make(chan int)
	go producer(chc)
	for v := range chc {
		fmt.Println("Received ", v)
	}

	// for range 循环信道示例
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	
	go calcSquares2(number, sqrch)
	go calcCubes2(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

// 实现获取一个数的每位数的逻辑
func digits(number int, ch chan int) {
	for number != 0 {
		digit := number % 10
		ch <- digit
		number /= 10
	}
	close(ch)
}

func calcSquares2(number int, sqrch chan int) {
	sum := 0
	ch := make(chan int)
	go digits(number, ch)
	for digit := range ch {
		sum += digit * digit
	}
	sqrch <- sum
}

func calcCubes2(number int, cubech chan int) {
	sum := 0
	ch := make(chan int)
	go digits(number, ch)
	for digit := range ch {
		sum += digit * digit * digit
	}
	cubech <- sum
}

func producer(chb chan int) {
	for i := 0; i < 10; i++ {
		chb <- i
	}
	close(chb)
}
