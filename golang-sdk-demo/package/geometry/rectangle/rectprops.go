package rectangle

import (
	"fmt"
	"math"
)

/*
所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。
func init() {
}

init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性。

包的初始化顺序如下：

首先初始化包级别（Package Level）的变量;
紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用。

如果一个包导入了另一个包，会先初始化被导入的包。
尽管一个包可能会被导入多次，但是它只会被初始化一次。
*/

func init() {
	fmt.Println("rectangle package initialized")
}

// Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
