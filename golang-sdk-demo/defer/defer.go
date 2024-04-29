package main

import "fmt"

/*
defer 语句的用途是：含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。
*/

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	// 延迟函数
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	/*输出：
	Started finding largest
	Largest number in [78 109 2 563 300] is 563
	Finished finding largest
	*/

	// 延迟方法
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ") //输出：Welcome John Smith

	/*
		实参取值（Arguments Evaluation）：
		在 Go 语言中，并非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值。
	*/
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

	/*输出：
	value of a before deferred function call 10
	value of a in deferred function 5
	*/

	/*
			defer 栈：
			当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，随后按照后进先出（Last In First Out, LIFO）的顺序执行。

		    示例:将一个字符串逆序打印。
	*/
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

	/*输出：
	Orignal String: Naveen
	Reversed String: neevaN
	*/
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
