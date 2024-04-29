package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)

		// 恢复 panic 时，我们就释放了它的堆栈跟踪, 使用下面语句实现恢复后获得堆栈跟踪
		debug.PrintStack()
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	// 运行时错误（如数组越界）也会导致 panic。
	a()
	fmt.Println("normally returned from main")

	/*输出：
	Recovered runtime error: index out of range [3] with length 3
	goroutine 1 [running]:
	runtime/debug.Stack()
	        /opt/go1.21.4/src/runtime/debug/stack.go:24 +0x5e
	runtime/debug.PrintStack()
	        /opt/go1.21.4/src/runtime/debug/stack.go:16 +0x13
	main.r()
	        /home/fdc/go/src/dcfenga/cloud-native-practice/golang-sdk-demo/panic-recover/panic-runtime.go:13 +0x70
	panic({0x48db20?, 0xc000114000?})
	        /opt/go1.21.4/src/runtime/panic.go:914 +0x21f
	main.a()
	        /home/fdc/go/src/dcfenga/cloud-native-practice/golang-sdk-demo/panic-recover/panic-runtime.go:20 +0x38
	main.main()
	        /home/fdc/go/src/dcfenga/cloud-native-practice/golang-sdk-demo/panic-recover/panic-runtime.go:26 +0x13
	normally returned from main
	*/
}
