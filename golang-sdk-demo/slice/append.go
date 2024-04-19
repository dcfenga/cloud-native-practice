package main

import "fmt"

/*
一个slice有三个基本组成元素：一个指向首元素的指针，一个长度值和一个容量值：

	type slice struct{
		Ptr *int //指向分配的数组的指针
		Len int  // 长度
		Cap int  // 容量
	}

参考：http://sharecore.net/post/%E5%AF%B9go%E7%9A%84slice%E8%BF%9B%E8%A1%8Cappend%E7%9A%84%E4%B8%80%E4%B8%AA%E5%9D%91/
*/
func main() {
	/*
		输出：append后，底层数组正好容纳append后的数据
		slice1: [2 6 7 8]
		arr1: [1 2 6 7 8]
	*/
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:2]
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)

	/*
		        分析：append后，当原slice容量足够，不需要进行扩容时，那对slice元素的追加，都是发生在原slice里的（数组里），所以，原数组被“悄悄”改变了。
			    元素个数会导致超出原slice的容量限制时，底层数组容纳不下append后的数据，此时执行下面步骤：
			    A.创建一个容量更大的slice（扩容）。与对slice进行切片操作不同，这个slice是全新的，它的数组也是全新的，指针也是指向新数组的首位置;
				B.新slice创建好后，会将原来被append的slice的元素内容进行值复制到新的slice;
				C.将要被append元素，追加到新slice的末尾；
				D.重建数据，所以原数组数据不变。
	*/
	/*
		输出：
		slice2: [2 3 6 7 8]
		arr2: [1 2 3 4 5]
	*/
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3]
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr2:", arr2)

	// 调用make方法来创建一个slice： func make([]T, len, [cap]) []T
	slicea := make([]int, 5, 5) // Go的默认零值：slicea == []int{0, 0, 0, 0, 0}
	fmt.Println(slicea)
	fmt.Println("-----------")

	// 调用make时，当cap参数未指定，那它的值与len相同
	sliceb := make([]int, 5)
	// len(sliceb)==5
	// cap(sliceb)==5
	fmt.Println(sliceb)
	fmt.Println("-----------")

	// 针对原slice或数组进行切片的方式来创建切片
	arra := [5]int{1, 2, 3, 4, 5}
	s1 := arra[1:4] //对数组进行切片
	fmt.Println(s1)
	fmt.Println("len(s1):", len(s1)) // len(s1)== 3 //len为切片开始位置到结束位置的个数
	fmt.Println("cap(s1):", cap(s1)) // cap(s1)=4 //容量为原数组总长度减开始位置

	s2 := s1[2:]                     //对slice:s1进行切片, 对slice进行切片操作，并不会新创建一个对象（分配内存），而只是在原来slice的基础上移动指针位置。
	fmt.Println("len(s2):", len(s2)) //len(s2)==1 //len为切片开始位置到结束位置的个数
	fmt.Println("cap(s2):", cap(s2)) //cap(s2)==2 //容量为原slice总容量减开始位置(4-2)
}
