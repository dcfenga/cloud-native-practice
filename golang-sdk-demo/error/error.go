package main

import (
	"fmt"
	"path/filepath"
)

/*
Golang中，错误用内建的 error 类型来表示。

按照 Go 的惯例，在处理错误时，通常都是将返回的错误与 nil 比较。nil 值表示了没有错误发生，而非 nil 值表示出现了错误。
*/
func main() {
	//f, err := os.Open("/README.md")
	//if err != nil {
	//	fmt.Println(err) // 输出：open /README.md: no such file or directory
	//	return
	//}
	//fmt.Println(f.Name(), "opened successfully")

	// 从错误获取更多信息的不同方法
	// 1. 断言底层结构体类型，使用结构体字段获取更多信息
	//if err, ok := err.(*os.PathError); ok {
	//	fmt.Println("File at path", err.Path, "failed to open") //输出：File at path /README.md failed to open
	//	return
	//}
	//fmt.Println(f.Name(), "opened successfully")

	// 2. 断言底层结构体类型，调用方法获取更多信息
	//addr, err := net.LookupHost("golangbot123.com")
	//if err, ok := err.(*net.DNSError); ok {
	//	if err.Timeout() {
	//		fmt.Println("operation timed out")
	//	} else if err.Temporary() {
	//		fmt.Println("temporary error")
	//	} else {
	//		fmt.Println("generic error: ", err)
	//	}
	//	return
	//}
	//fmt.Println(addr)
	// 输出：generic error:  lookup golangbot123.com on 10.7.20.241:53: no such host

	// 3. 直接比较
	//files, error := filepath.Glob("[")
	//if error != nil && error == filepath.ErrBadPattern {
	//	fmt.Println(error) // 输出： syntax error in pattern
	//	return
	//}
	//fmt.Println("matched files", files)

	// 不可忽略错误
	// 通过使用 _ 空白标识符，忽略 Glob 函数返回的错误
	files2, _ := filepath.Glob("[")
	fmt.Println("matched files", files2) // 输出：matched files []
}
