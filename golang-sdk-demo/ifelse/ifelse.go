package main

import "fmt"

func main() {
	// 单分支判断
	age := 28
	gender := "male"

	// if 使用方式1，只有if condition {}
	// &&：表示且，左右都需要为true，最终结果才能为 true，否则为 false
	// ||：表示或，左右只要有一个为true，最终结果即为true，否则 为 false
	if age > 20 && gender == "male" {
		fmt.Println("是成年男性")
	}
	fmt.Println("-------------")

	// if 使用方式2, if statement; condition {}
	// 即在 if 里可以允许先运行一个表达式，取得变量后，再对其进行判断
	// 但此时变量只能在if和else中访问，从if和else外访问编译器不通过
	if age2 := 20; age2 > 18 {
		fmt.Println("已经成年了")
	}
	fmt.Println("-------------")

	// if 使用方式3，多分支判断
	if age > 18 {
		fmt.Println("已经成年了")
	} else if age > 12 {
		fmt.Println("已经是青少年了")
	} else {
		fmt.Println("还不是青少年")
	}

	// if 使用方式注意点：else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	}
	// else 处异常：syntax error: unexpected else, expected }，原因是Go 语言的分号是自动插入的，
	// 即在 Go 语言规则中，它指定在 } 之后插入一个分号，如果这是该行的最终标记。因此，在 if 语句后面的 } 会自动插入一个分号。
	//else {
	//	fmt.Println("the number is odd")
	//}
}
