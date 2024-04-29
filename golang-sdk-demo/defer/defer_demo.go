package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

/*
优化：wg.Done() 只在 area 函数返回的时候才会调用。wg.Done() 应该在 area 将要返回之前调用，
并且与代码流的路径（Path）无关，因此我们可以只调用一次 defer，来有效地替换掉 wg.Done() 的多次调用。
*/
func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	/*输出：
	rect {8 9}'s area 72
	rect {-67 89}'s length should be greater than zero
	rect {5 -67}'s width should be greater than zero
	All go routines finished executing
	*/
}
