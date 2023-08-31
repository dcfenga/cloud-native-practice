package main

import (
	"fmt"
	"math"
)

/*
FMT格式化：
%b    表示为二进制
%c    该值对应的unicode码值
%d    表示为十进制
%o    表示为八进制
%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x    表示为十六进制，使用a-f
%X    表示为十六进制，使用A-F
%U    表示为Unicode格式：U+1234，等价于"U+%04X"
*/
func main() {
	// 2进制：以0b或0B为前缀
	var num01 int = 0b1100

	// 8进制：以0o或者 0O为前缀
	var num02 int = 0o14

	// 16进制：以0x 为前缀
	var num03 int = 0xC

	fmt.Println(num01, num02, num03)
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	// 浮点数精度主要取决于尾数部分的位数。
	// 对于 float32（单精度）来说，表示尾数的为23位，除去全部为0的情况以外，最小为2^-23，约等于1.19*10^-7，
	// 所以float小数部分只能精确到后面6位，加上小数点前的一位，即有效数字为7位。
	//
	// 同理 float64（单精度）的尾数部分为 52位，最小为2^-52，约为2.22*10^-16，
	// 所以精确到小数点后15位，加上小数点前的一位，有效位数为16位。
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	// 精度测试，临界情况
	// 参考: https://www.zhihu.com/question/26022206
	var myfloat float32 = 10000018
	fmt.Println("myfloat: ", myfloat)
	fmt.Println("myfloat: ", myfloat+1)

	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187
	fmt.Println("myfloat01: ", myfloat01)
	fmt.Println("myfloat02: ", myfloat01+5)
	fmt.Println(myfloat02 == myfloat01+5)
}
