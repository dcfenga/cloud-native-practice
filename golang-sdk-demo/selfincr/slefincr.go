package main

import "fmt"

func main() {
	// Go 只支持后自增/自减，不支持前缀自增/自减
	a := 1
	a++
	fmt.Println(a)
	fmt.Println("------------------")

	// Go中++是语句，不是表达式，即不可以作为表达式被赋值给其它的变量使用
	// 下面异常：syntax error: unexpected ++ at end of statement
	// b := a++

	c := a
	fmt.Println(c)

	// Go中不支持前缀++/--,下面异常：syntax error: unexpected ++, expected }
	// ++a
}
