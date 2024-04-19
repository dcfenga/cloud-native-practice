package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Go 语言中的字符串是一个字节切片,需把内容放在双引号""之间；
Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。
*/
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		// %x 格式限定符用于指定 16 进制编码, 输出：48 65 6c 6c 6f 20 57 6f 72 6c 64
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		// %c 格式限定符用于打印字符串的字符, 输出：H e l l o   W o r l d
		// 存在BUG
		fmt.Printf("%c ", s[i])
	}
}

func printCharsRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
		// %c 格式限定符用于打印字符串的字符, 输出：S e ñ o r
		fmt.Printf("%c ", runes[i])
	}
}

/*
返回的是是当前 rune 的字节位置。程序的输出结果为：

S starts at byte 0
e starts at byte 1
ñ starts at byte 2
o starts at byte 4
r starts at byte 5

从上面的输出中可以清晰的看到 ñ 占了两个字节。
*/
func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Hello World"
	fmt.Println(name)

	// 单独获取字符串的每一个字节
	printBytes(name)
	fmt.Printf("\n")

	printChars(name)
	fmt.Printf("\n")

	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")

	// 异常，输出：S e Ã ± o r
	// 分析：因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。它的 UTF-8 编码[7]占用了 c3 和 b1 两个字节。
	// 它的 UTF-8 编码占用了两个字节 c3 和 b1。而我们打印字符时，却假定每个字符的编码只会占用一个字节，这是错误的。
	// 在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。
	// 采用rune类型解决这个问题。
	printChars(name)
	fmt.Printf("\n")

	// rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
	printCharsRune(name)
	fmt.Printf("\n")

	// 字符串的 for range 循环
	name2 := "Señor"
	printCharsAndBytes(name2)
	fmt.Printf("\n")

	// 用字节切片构造字符串
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str) //输出：Café
	fmt.Printf("\n")

	byteSlice2 := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str2 := string(byteSlice2)
	fmt.Println(str2) //输出：Café
	fmt.Printf("\n")

	// 用 rune 切片构造字符串
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str3 := string(runeSlice)
	fmt.Println(str3) //输出：Señor
	fmt.Printf("\n")

	// 字符串的长度、
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)
	fmt.Printf("\n")

	/*输出：
	length of Señor is 5 by utf8.RuneCountInString(s)
	length of Señor is 6 by len(s)
	length of Pets is 4 by utf8.RuneCountInString(s)
	length of Pets is 4 by len(s)
	*/

	// 字符串是不可变的
	// Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改。
	s := "hello"
	fmt.Println(mutate(s)) //输出：aello

	fmt.Println(mutate2([]rune(s))) //输出：aello
}

func mutate(s string) string {
	// 异常：cannot assign to s[0] (neither addressable nor a map index expression)
	// s[0] = 'a'

	runes := []rune(s)
	runes[0] = 'a'
	return string(runes) //返回：aello
}

func mutate2(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func length(s string) {
	// utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度
	fmt.Printf("length of %s is %d by utf8.RuneCountInString(s)\n", s, utf8.RuneCountInString(s))
	fmt.Printf("length of %s is %d by len(s)\n", s, len(s))
}
