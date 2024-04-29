package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

// 优化后仅在area()返回时调用一次wg.done()，但在if语句中增加return
func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
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
