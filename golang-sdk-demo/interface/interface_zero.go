package main

import "fmt"

/*
接口零值：
接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
*/

type Describer interface {
	Describe()
}

func main() {
	var d Describer
	if d == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d, d)
		// 输出：d1 is nil and has type <nil> value <nil>
	}
	d.Describe()
	// 异常：panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x47c53b]
	// 分析：d 等于nil
}
