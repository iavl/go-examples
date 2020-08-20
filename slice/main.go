package main

import (
	"fmt"
	"unsafe"
)

func a() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)
	fmt.Println(y, z)
}

func b() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}

func f(arr [][]int) {
	fmt.Println(fmt.Sprintf("%p", arr))

}

func ap(a []int) {
	//a[0] = 1
	a = append(a, 10)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)

}

func main() {
	//a()
	//b()

	//var arr = [][]int{
	//	{1, 1, 1},
	//	{1, 0, 1},
	//	{1, 1, 1},
	//}
	//fmt.Println(fmt.Sprintf("%p", arr))
	//f(arr)

	//a := []int{7, 8, 9}
	//fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	//ap(a)
	//fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)

	a := []int{}
	a = append(a, 7, 8, 9)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	ap(a)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	p := unsafe.Pointer(&a[2])
	q := uintptr(p) + 8
	t := (*int)(unsafe.Pointer(q))
	fmt.Println(*t)
}
