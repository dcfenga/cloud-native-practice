package main

import (
	"fmt"
	"math"
)

func main() {
	// 浮点数精度主要取决于尾数部分的位数。
	// 对于 float32（单精度）来说，表示尾数的为23位，除去全部为0的情况以外，最小为2^-23，约等于1.19*10^-7，
	// 所以float小数部分只能精确到后面6位，加上小数点前的一位，即有效数字为7位。
	//
	// 同理 float64（单精度）的尾数部分为 52位，最小为2^-52，约为2.22*10^-16，
	// 所以精确到小数点后15位，加上小数点前的一位，有效位数为16位。
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

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
