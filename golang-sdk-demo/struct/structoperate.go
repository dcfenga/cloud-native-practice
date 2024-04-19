package main

import "fmt"

type Employee3 struct {
	firstName, lastName string
	age, salary         int
}

// 定义一个与 Employee3 的绑定的方法, 其中 FmtEmployee 是方法名，而(person Employee3)：
// 表示将 FmtEmployee 方法与 Employee3 的实例绑定。把 Employee3 称为方法的接收者，
// 而 person 表示实例本身，它相当于 Python 中的 self，在方法内可以使用
// person.属性名 的方法来访问实例属性。
// 以值做为方法接收者
func (person Employee3) FmtEmployee() {
	fmt.Printf("First Name： %s\n", person.firstName)
	fmt.Printf("Last Name： %s\n", person.lastName)
	fmt.Printf("Age： %d\n", person.age)
	fmt.Printf("Salary： %d\n", person.salary)
}

// 以指针做为方法接收者, 下面几种情况应当直接使用指针做为方法的接收者：
// 1.需要在方法内部改变结构体内容的时候
// 2.考虑性能的问题，当结构体过大的时候
// 3.虑到代码一致性，建议都使用指针做为接收者
func (person *Employee3) increase_Age() {
	person.age += 1
}

func main() {
	// 访问结构体的字段: 点号操作符 . 用于访问结构体的字段
	emp6 := Employee3{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
	fmt.Println()

	// 创建零值的 struct，以后再给各个字段赋值
	var emp7 Employee3
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee3 7:", emp7)

	// 创建指向结构体的指针
	// 可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName
	emp8 := &Employee3{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
	fmt.Println("=================")
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	// 实例化
	myself := Employee3{firstName: "小明", age: 30, salary: 10000}

	// 调用函数
	myself.FmtEmployee()

	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.increase_Age()
	fmt.Printf("当前年龄：%d\n", myself.age)
}
