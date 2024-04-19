package main

import "fmt"

/*
for 循环语法:
for initialisation; condition; post {
}
初始化语句只执行一次。循环初始化后，将检查循环条件。如果条件的计算结果为 true ，则 {} 内的循环体将执行，
接着执行 post 语句。post 语句将在每次成功循环迭代后执行。在执行 post 语句后，条件将被再次检查。
如果为 true, 则循环将继续执行, 否则 for 循环将终止。
*/
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println("\n-----------------")

	// break 语句用于在完成正常执行之前突然终止 for 循环，之后程序将会在 for 循环下一行代码开始执行。
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
	fmt.Println("\n-----------------")

	// continue 语句用来跳出 for 循环中当前循环。在 continue 语句后的所有的 for 循环语句都不会在本次循环中执行。循环体会在一下次循环中继续执行。
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n-----------------")

	// 简化for
	i := 0
	for i <= 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println("\n-----------------")

	// for 循环中声明和操作多个变量
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//无限循环的语法是：
	//for {
	//}
}
