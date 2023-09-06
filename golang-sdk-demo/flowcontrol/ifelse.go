package main

import "fmt"

func main() {
	// 单分支判断
	age := 28
	gender := "male"

	// &&：表示且，左右都需要为true，最终结果才能为 true，否则为 false
	// ||：表示或，左右只要有一个为true，最终结果即为true，否则 为 false
	if age > 20 && gender == "male" {
		fmt.Println("是成年男性")
	}
	fmt.Println("-------------")

	// 简写
	// 即在 if 里可以允许先运行一个表达式，取得变量后，再对其进行判断
	if age2 := 20; age2 > 18 {
		fmt.Println("已经成年了")
	}
	fmt.Println("-------------")

	// 多分支判断
	if age > 18 {
		fmt.Println("已经成年了")
	} else if age > 12 {
		fmt.Println("已经是青少年了")
	} else {
		fmt.Println("还不是青少年")
	}
}
