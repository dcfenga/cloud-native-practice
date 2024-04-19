package main

import (
	_ "dcfenga/cloud-native-practice/golang-sdk-demo/package/geometry/rectangle" //实现导入此包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。
	"fmt"
	"log"
)

/*
main 包的初始化顺序为：

首先初始化被导入的包。因此，首先初始化了 rectangle 包;
接着初始化了包级别的变量 rectLen 和 rectWidth;
调用 init 函数;
最后调用 main 函数。
*/

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

/*
错误屏蔽器,用于程序开发的活跃阶段，又常常会先导入包，而暂不使用它。
*/
//var _ = rectangle.Area(rectLen, rectWidth)

func main() {
	fmt.Println("Geometrical shape properties")

	/*Area function of rectangle package used*/
	//fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	//fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
