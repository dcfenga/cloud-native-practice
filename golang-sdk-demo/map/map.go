package main

import "fmt"

/*
map 是在 Go 中将值（value）与键（key）关联的内置类型。
通过向 make 函数传入键和值的类型，可以创建 map，创建map语法：
personSalary := make(map[string]int)

其中keys只能是可比较的类型，如 boolean，interger，float，complex，string 等，都可以作为键
参考: http://golang.org/ref/spec#Comparison_operators
*/
func main() {
	//创建 map: map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)

		// 给 map 添加元素
		personSalary["steve"] = 12000
		personSalary["jamie"] = 15000
		personSalary["mike"] = 9000
		fmt.Println("personSalary map contents:", personSalary)
	}
	fmt.Println("===============")

	// 声明的时候初始化 map
	personSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary2["mike"] = 9000
	fmt.Println("personSalary2 map contents:", personSalary2)
	fmt.Println("===============")

	// 获取 map 中的元素
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])

	// 获取一个不存在的元素，返回value类型的零值
	fmt.Println("Salary of joe is", personSalary["joe"])
	fmt.Println("===============")

	// 判断key是否存在，语法：value, ok := map[key]
	newEmp := "joe"
	value, ok := personSalary2[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}
	fmt.Println("===============")

	// 遍历 map 中所有的元素需要用 for range 循环
	// 不保证每次执行程序获取的元素顺序相同
	fmt.Println("All items of a map")
	for key, value := range personSalary2 {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
	fmt.Println("===============")

	// 删除 map 中的元素
	fmt.Println("map before deletion", personSalary2)
	delete(personSalary2, "steve")
	fmt.Println("map after deletion", personSalary2)
	fmt.Println("===============")

	// 获取 map 的长度
	fmt.Println("length is", len(personSalary))
	fmt.Println("===============")

	// Map 是引用类型
	// 和 slices类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构;
	// 因此，改变其中一个变量，就会影响到另一变量。
	// 当 map 作为函数参数传递时也会发生同样的情况。函数中对 map 的任何修改，对于外部的调用都是可见的。
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
	fmt.Println("===============")

	// Map 的相等性
	// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
	// 判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素。
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	// 异常：invalid operation: map1 == map2 (map can only be compared to nil)
	//if map1 == map2 {
	//}
	fmt.Println("map2 map contents:", map2)
}
