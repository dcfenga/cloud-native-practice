package main

import "fmt"

/*
1. 声明结构体

	type 结构体名 struct {
	    属性名   属性类型
	    属性名   属性类型
	    ...
	}

2. 定义方法
使用组合函数的方式来定义结构体方法

3. 方法可见性
(1)当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用;
(2)当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
*/

// Profile 定义一个名为Profile 的结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// FmtProfile 定义一个与 Profile 的绑定的方法, 其中FmtProfile 是方法名，而(person Profile)：
// 表示将 FmtProfile 方法与 Profile 的实例绑定。把 Profile 称为方法的接收者，
// 而 person 表示实例本身，它相当于 Python 中的 self，在方法内可以使用
// person.属性名 的方法来访问实例属性。
// 以值做为方法接收者
func (person Profile) FmtProfile() {
	fmt.Printf("名字： %s\n", person.name)
	fmt.Printf("年龄： %d\n", person.age)
	fmt.Printf("性别： %s\n", person.gender)
}

// 以指针做为方法接收者, 下面几种情况应当直接使用指针做为方法的接收者：
// 1.你需要在方法内部改变结构体内容的时候
// 2.出于性能的问题，当结构体过大的时候
// 3.虑到代码一致性，建议都使用指针做为接收者
func (person *Profile) increase_Age() {
	person.age += 1
}

func main() {
	// 实例化
	myself := Profile{name: "小明", age: 30, gender: "male"}

	// 调用函数
	myself.FmtProfile()

	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.increase_Age()
	fmt.Printf("当前年龄：%d", myself.age)
}
