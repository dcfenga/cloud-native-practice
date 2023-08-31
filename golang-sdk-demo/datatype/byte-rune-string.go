package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// byte，占用1个节字，就 8 个比特位，所以它和 uint8 类型本质上没有区别，
	// 它表示的是 ACSII 表中的一个字符。
	var a byte = 65
	// 8进制写法: var c byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var c byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c \nb 的值: %c \n", a, b)

	// 或者使用 string 函数
	fmt.Println("a 的值: ", string(a), " \nb 的值: ", string(b))

	var a2 byte = 'A'
	var b2 uint8 = 'B'
	fmt.Printf("a2 的值: %c \nb2 的值: %c \n", a2, b2)

	// rune，占用4个字节，共32位比特位，所以它和 uint32 本质上也没有区别。
	// 它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范）。
	var a3 byte = 'A'
	var b3 rune = 'B'
	fmt.Printf("a3 占用 %d 个字节数\nb3 占用 %d 个字节数 \n", unsafe.Sizeof(a3), unsafe.Sizeof(b3))

	// 单引号表示字符，双引号表示字符串
	// 因为uint8 和 uint32 ，直观上让人以为这是一个数值，但是实际上，它也可以表示一个字符，
	// 所以为了消除这种直观错觉，就诞生了 byte 和 rune 这两个别名类型。

	var mystr1 string = "hello"
	var mystr2 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr1: %s\n", mystr1)
	fmt.Printf("mystr2: %s\n", mystr2)

	// 英文字母占用一个字节，而中文字母占用 3个字节
	var country string = "hello,中国"
	fmt.Println(len(country))
}
