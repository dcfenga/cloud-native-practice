package main

import "fmt"

/*
Go 不支持继承，但它支持组合（Composition）。组合一般定义为“合并在一起”。
汽车就是一个关于组合的例子：一辆汽车由车轮、引擎和其他各种部件组合在一起。
*/

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)

	// 一旦结构体内嵌套了一个结构体字段，Go 可以使我们访问其嵌套的字段，好像这些字段属于外部结构体一样
	fmt.Println("Author: ", p.author.fullName()) // 可替换为：p.fullName()
	fmt.Println("Bio: ", p.author.bio)           // 可替换为：p.bio()
}

type website struct {
	// 结构体不能嵌套一个匿名切片，必须有名字
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	// 通过嵌套结构体进行组合
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}

	post1.details()

	/*输出：
	Title:  Inheritance in Go
	Content:  Go supports composition instead of inheritance
	Author:  Naveen Ramanathan
	Bio:  Golang Enthusiast
	*/

	//结构体切片的嵌套
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()

	/*输出：
	Title:  Inheritance in Go
	Content:  Go supports composition instead of inheritance
	Author:  Naveen Ramanathan
	Bio:  Golang Enthusiast
	Contents of Website

	Title:  Inheritance in Go
	Content:  Go supports composition instead of inheritance
	Author:  Naveen Ramanathan
	Bio:  Golang Enthusiast

	Title:  Struct instead of Classes in Go
	Content:  Go does not support classes but methods can be added to structs
	Author:  Naveen Ramanathan
	Bio:  Golang Enthusiast

	Title:  Concurrency
	Content:  Go is a concurrent language and not a parallel one
	Author:  Naveen Ramanathan
	Bio:  Golang Enthusiast
	*/
}
