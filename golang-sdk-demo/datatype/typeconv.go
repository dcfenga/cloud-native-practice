package main

import "fmt"

/*
 * Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。
 */
func main() {
	i := 55   //int
	j := 67.8 //float64
	// sum := i + j //不允许 int + float64, 异常： invalid operation: i + j (mismatched types int and float64)
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)

	m := 100
	var n float64 = float64(m) // 若没有显式转换，该语句会报错
	fmt.Println("n", n)
}
