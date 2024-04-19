package main

import (
	"dcfenga/cloud-native-practice/golang-sdk-demo/struct/computer"
	"fmt"
)

/*
如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。
同样，如果结构体里的字段首字母大写，它也能被其他包访问到。
*/
func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	// spec.model = "Mac Mini" //异常：spec.model undefined (type computer.Spec has no field or method model)
	fmt.Println("Spec:", spec)
}
