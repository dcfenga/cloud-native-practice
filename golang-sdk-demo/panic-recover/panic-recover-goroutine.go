package main

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recoverd:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("on! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")

	/*输出：
	Inside A
	Inside B
	panic: on! B panicked

	goroutine 18 [running]:
	main.b()
	        /home/fdc/go/src/dcfenga/cloud-native-practice/golang-sdk-demo/panic-recover/panic-recover-goroutine.go:23 +0x5f
	created by main.a in goroutine 1
	        /home/fdc/go/src/dcfenga/cloud-native-practice/golang-sdk-demo/panic-recover/panic-recover-goroutine.go:17 +0x79

	分析：调用 recovery 的协程和 b() 中发生 panic 的协程并不相同，因此不可能恢复 panic。
	*/
}
