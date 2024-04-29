package main

import (
	"fmt"
	"reflect"
)

/*
反射就是程序能够在运行时检查变量和值，求出它们的类型。

在 Go 语言中，`reflect`实现了运行时反射。reflect 包会帮助识别 `interface{}`变量的底层具体类型和具体值。

reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的具体值。reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别返回 reflect.Type 和 reflect.Value。
*/
type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func createQuery2(q interface{}) {
	t := reflect.TypeOf(q)  // Type 表示 interface{} 的实际类型（在这里是 main.Order)
	v := reflect.ValueOf(q) // Kind 表示该类型的特定类别（在这里是 struct）
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)
}

func createQuey3(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField()) // NumField()`方法返回结构体中字段的数量
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i)) // Field(i int)`方法返回字段 i 的 reflect.Value
		}
	}
}

// createQuery4 函数应该适用于所有的结构体。因此，要编写这个函数，就必须在运行时检查传递过来的结构体参数的类型，
// 找到结构体字段，接着创建查询语句。
func createQuery4(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name() // 从结构体的 reflect.Type 获取结构体的名字
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	// 运行时检查变量，确定变量类型
	i := 10 // i 的类型在编译时就已确定
	fmt.Printf("%d %T\n", i, i)

	// 示例：接收结构体作为参数，并用它来创建一个 SQL 插入查询
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o)) // 输出：insert into order values(1234, 567)

	createQuery2(o)
	/*输出：
	Type  main.order
	Value  {1234 567}
	Kind  struct
	*/

	createQuey3(o)
	/*输出：
	Number of fields 2
	Field:0 type:reflect.Value value:1234
	Field:1 type:reflect.Value value:567
	*/

	// Int() 和 String() 方法用于分别取出 reflect.Value 作为 int64 和 string
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

	// 示例：实现根据结构体输出SQL语句
	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery4(e) // 输出：insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
	z := 90
	createQuery4(z) // 输出：unsupported type
}
