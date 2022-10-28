package data

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 切片 slice 的数据结构#
// type SliceHeader struct {
//     Data uintptr  //指向的底层数组地址
//     Len  int     //切片长度
//     Cap  int     //切片容量
// }

func SliceCopy() {
	// 浅 copy
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	fmt.Println(slice1) // [1 2 3 4 5]
	fmt.Println(slice2) // [1 2 3 4 5]
	//同时改变两个数组
	slice1[1] = 100
	fmt.Println(slice1) // [1 100 3 4 5]
	fmt.Println(slice2) // [1 100 3 4 5]
	fmt.Println("切片1指向的底层数组地址:", (*reflect.SliceHeader)(unsafe.Pointer(&slice1)))
	fmt.Println("切片2指向的底层数组地址:", (*reflect.SliceHeader)(unsafe.Pointer(&slice2)))

	// 深 copy
	slice01 := []int{1, 2, 3, 4, 5}
	slice02 := make([]int, 5)
	copy(slice02, slice01)
	fmt.Println(slice01)
	fmt.Println(slice02)
	slice1[1] = 100 // 只会影响slice1
	fmt.Println(slice01)
	fmt.Println(slice02)
	fmt.Println("切片1指向的底层数组地址:", (*reflect.SliceHeader)(unsafe.Pointer(&slice1)))
	fmt.Println("切片2指向的底层数组地址:", (*reflect.SliceHeader)(unsafe.Pointer(&slice2)))
}
