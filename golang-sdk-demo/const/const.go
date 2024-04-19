package main

import (
	"fmt"
	"math"
)

/*
 * Go 语言中，术语"常量"用于表示固定的值。
 */

func main() {
	// 常量不能再重新赋值为其他的值
	const x = 5 // 允许
	// x = 80 //不允许重新赋值，异常： Cannot assign to x

	// 常量的值会在编译的时候确定。因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。
	fmt.Println("Hello, playground")
	var m = math.Sqrt(4) // 允许
	//const n = math.Sqrt(4) // 不允许,异常：Const initializer 'math.Sqrt(4)' is not a constant
	fmt.Println("m", m)
	// fmt.Println("n", n)

	// 无类型的字符串属于常量,Go 是一门强类型语言，所有的变量必须有明确的类型;
	// 所以无类型的常量有一个与它们相关联的默认类型，并且当且仅当一行代码需要时才提供它。
	const hello = "Hello World"
	var name = "Sam"
	fmt.Printf("\ntype %T value %v", hello, hello)
	fmt.Printf("\ntype %T value %v", name, name)
	fmt.Println("\n===============================")

	// 创建一个带类型的常量
	const typehello string = "hello World" // typedhello 就是一个 string 类型的常量。
	var defaultName = "Sam"                // 允许
	type myString string                   // 新类型 myString，它是 string 的别名。
	var customName myString = "Sam"        // 允许,常量 Sam 是无类型的，它可以分配给任何字符串变量
	// 不允许,Go 的类型策略不允许将一种类型的变量赋值给另一种类型的变量。
	// 因此将 defaultName 赋值给 customName 是不允许的。
	// 异常：cannot use defaultName (variable of type string) as myString value in assignment
	// customName = defaultName
	fmt.Printf("\ntype %T value %v", typehello, typehello)
	fmt.Printf("\ntype %T value %v", defaultName, defaultName)
	fmt.Printf("\ntype %T value %v", customName, customName)
	fmt.Println("\n===============================")

	// 布尔常量
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       // 允许
	var customBool myBool = trueConst // 允许
	// defaultBool = customBool          // 不允许,异常：Cannot use 'customBool' (type myBool) as the type bool
	fmt.Printf("\ntype %T value %v", trueConst, trueConst)
	fmt.Printf("\ntype %T value %v", defaultBool, defaultBool)
	fmt.Printf("\ntype %T value %v", customBool, customBool)
	fmt.Println("\n===============================")

	// 数字常量
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
	fmt.Println("\n===============================")

	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// 数字表达式
	var y = 5.9 / 8
	fmt.Printf("y's type %T value %v", y, y)
}
