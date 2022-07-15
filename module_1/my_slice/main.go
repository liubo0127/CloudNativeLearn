package main

import "fmt"

func setValue(s []int) {
	a := []int{1, 2, 3} // a 长度为 3
	c := a              // c 长度也为 3
	fmt.Printf("a address: %p %p %p\n", &a[0], &a[1], &a[2])
	fmt.Printf("c address: %p %p %p\n", &c[0], &c[1], &c[2])
	a = append(s, 1) // 录播说法：当 s 长度 >= 3 时，a 的长度不足以存储新的切片，会分配新的内存来存储，，此时 a != c。实测下来不管 b 的长度是否会超过 a 的长度，都会重新分配内存地址
	fmt.Printf("a address: %p %p %p\n", &a[0], &a[1], &a[2])
	fmt.Printf("c address: %p %p %p\n", &c[0], &c[1], &c[2])
}

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
	mySlice1 := array1[1:3] // 数组连续片段的引用，这里的范围是左闭右开 [2,3]
	mySlice2 := array1[1:]  // 从下标为 1 的元素一直到最后一个元素构成的切片
	mySlice3 := array1[:3]  // 从下标为 0 的元素开始到下标为 2 的元素构成的切片
	var mySlice4 []int
	mySlice4 = append(mySlice4, 1) // mySlice2 中添加元素
	mySlice4 = append(mySlice4, 2)
	fmt.Println(mySlice1, mySlice2, mySlice3, mySlice4)
	s := []int{1, 2}
	//fmt.Printf("s address: %p %p\n", &s[0], &s[1])
	setValue(s)
}
