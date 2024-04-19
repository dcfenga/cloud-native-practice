package main

import "fmt"

func main() {
	// Go 中，真值用 true 表示，不但不与 1 相等，并且更加严格，不同类型无法进行比较，而假值用 false 表示，同样与 0 无法比较。
	var male bool = true
	// fmt.Println(male == 0) //IDEA中报错: Invalid operation: male == 0 (mismatched types bool and untyped int)
	fmt.Println(!male == false)
	fmt.Println(male != false)
	fmt.Println("----------------")

	// Go 语言中，则使用 && 表示且，用 || 表示或，并且有短路行为（即左边表达式已经可以确认整个表达式的值，那么右边将不会再被求值。
	var age int = 15
	var gender string = "male"
	//  && 两边的表达式都会执行
	fmt.Println(age > 18 && gender == "male")
	fmt.Println("****************")

	// gender == "male" 并不会执行
	fmt.Println(age > 14 || gender == "male")
	fmt.Println("----------------")

	// gender == "male" 会被执行
	fmt.Println(age > 16 || gender == "male")
	fmt.Println("----------------")
}

// bool转int
func bool2int(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

// int转bool
func int2bool(i int) bool {
	return i != 0
}
