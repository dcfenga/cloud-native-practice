package main

import (
	"fmt"
	"unsafe"
)

/*
Go基本类型：
  - bool
    表示一个布尔值，值为 true 或者 false
  - 数字类型
    int8: 表示 8 位有符号整型大小：8 位范围：-128～127
    int16: 表示 16 位有符号整型大小：16 位范围：-32768～32767
    int32: 表示 32 位有符号整型大小：32 位范围：-2147483648～2147483647
    int64: 表示 64 位有符号整型大小：64 位范围：-9223372036854775808～9223372036854775807
    int: 根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型
    uint8: 表示 8 位无符号整型大小：8 位范围：0～255
    uint16: 表示 16 位无符号整型大小：16 位范围：0～65535
    uint32: 表示 32 位无符号整型大小：32 位范围：0～4294967295
    uint64: 表示 64 位无符号整型大小：64 位范围：0～18446744073709551615
    uint: 根据不同的底层平台，表示 32 或 64 位无符号整型
    float32: 32 位浮点数
    float64: 64 位浮点数
    complex64: 实部和虚部都是 float32 类型的的复数
    complex128: 实部和虚部都是 float64 类型的的复数
    byte: uint8 的别名
    rune: int32 的别名
  - string: 字节的集合,即一个字符串就是由很多字符组成

FMT格式化：
%b    表示为二进制
%c    该值对应的unicode码值
%d    表示为十进制
%o    表示为八进制
%q    该值对应的单引号括起来的Go语法字符字面值，必要时会采用安全的转义表示
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

	// int 类型的大小在 32 位系统下是 32 位，而在 64 位系统下是 64 位
	var num04 int = 98
	num05 := 100
	fmt.Println("value of num04 is", num04, "and num05 is", num04)
	fmt.Printf("type of num04 is %T, size of num04 is %d", num04, unsafe.Sizeof(num04))   // num04 的类型和大小
	fmt.Printf("\ntype of num05 is %T, size of num05 is %d", num05, unsafe.Sizeof(num05)) // num05 的类型和大小
}
