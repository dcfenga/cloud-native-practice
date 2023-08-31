package main

import "fmt"

func main() {
	// map的key 不能是切片，不能是字典，不能是函数

	// 声明初始化字典方法1：声明时赋值
	var scoresa map[string]int = map[string]int{"english": 80, "chinese": 85}
	fmt.Println(scoresa)
	fmt.Println("----------------")

	// 声明初始化字典方法2：简化声明
	scoresb := map[string]int{"english": 80, "chinese": 85}
	fmt.Println(scoresb)
	fmt.Println("----------------")

	// 声明初始化字典方法3：通过make创建
	scoresc := make(map[string]int)
	scoresc["english"] = 80
	scoresc["chinese"] = 85
	fmt.Println(scoresc)
	fmt.Println("----------------")

	// 声明一个名为 score 的字典
	var scores map[string]int

	// 未初始化的 score 的零值为nil，无法直接进行赋值
	if scores == nil {
		// 需要使用 make 函数先对其初始化
		scores = make(map[string]int)
	}

	// 经过初始化后，就可以直接赋值
	scores["chinese"] = 90
	fmt.Println(scores)
	fmt.Println("----------------")

	// 添加元素
	scores["math"] = 95
	// 更性元素
	scores["chinese"] = 98
	fmt.Println(scores)
	fmt.Println(scores["math"])
	fmt.Println("----------------")

	// 删除元素
	delete(scores, "math")
	fmt.Println(scores)
	fmt.Println("----------------")

	// 获取默认值
	fmt.Println(scores["english"])
	fmt.Println("----------------")

	// 判断 key
	// 当key不存在，会返回value-type的零值 ，
	// 所以你不能通过返回的结果是否是零值来判断对应的 key 是否存在，因为 key 对应的 value 值可能恰好就是零值。
	// 字典的下标读取可以返回两个值，使用第二个返回值表示对应的 key 是否存在，若存在ok为true，若不存在，则ok为false。
	scoresx := map[string]int{"english": 80, "chinese": 85}
	if math, ok := scoresx["math"]; ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}
	fmt.Println("----------------")

	// map循环1：获取 key 和 value
	scoresy := map[string]int{"english": 80, "chinese": 85}
	for subject, score := range scoresy {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}
	fmt.Println("----------------")

	// map循环2: 只获取key，这里注意不用占用符
	for subject := range scoresy {
		fmt.Printf("key : %s\n", subject)
	}
	fmt.Println("----------------")

	// map循环3: 只获取 value，用一个占位符替代
	for _, score := range scoresy {
		fmt.Printf("value :%d\n", score)
	}
	fmt.Println("----------------")

	// Go 中，真值用 true 表示，不但不与 1 相等，并且更加严格，不同类型无法进行比较，而假值用 false 表示，同样与 0 无法比较。
	var male bool = true
	// fmt.Println(male == 0)   //IDEA中报错
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
